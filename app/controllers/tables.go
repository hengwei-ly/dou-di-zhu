package controllers

import (
	"doudizhu/app"

	"doudizhu/app/models"

	"github.com/revel/revel"
)

type Tables struct {
	App
}

const (
	StatusRedoubleSuccess = iota + 1
	StatusRedoubleRestart
	StatusRedoubleRecycle
)

func (c Tables) Index(tableId, seatId int) revel.Result {
	table := app.Tables[tableId]

	//对应的座位信息
	rightSeatId := (seatId+1)%3 + 1
	leftSeatId := (seatId+2)%3 + 1
	c.ViewArgs["mainSeatId"] = seatId
	c.ViewArgs["rightSeatId"] = rightSeatId
	c.ViewArgs["leftSeatId"] = leftSeatId

	return c.Render(table)
}

func (c Tables) ChangeTableStatus(tableId int) revel.Result {
	return c.RenderJSON(app.Tables[tableId])
}

func (c Tables) IntoTable(tableId, userId int) revel.Result {
	table := app.Tables[tableId]
	isFull := true
	for k := 0; k < len(table.Seats); k++ {
		if table.Seats[k].UserId == 0 {
			table.Seats[k].UserId = userId
			isFull = false
		}
	}

	var result = struct {
		SeatId  int
		IsError bool
		ErrMsg  string
	}{}

	if isFull {
		result.IsError = true
		result.ErrMsg = "该房间已满"
	}
	return c.RenderJSON(result)
}

//func (c DouDiZhus) Prepare(tableId, userId int) revel.Result {
//	return nil
//}

//没有确定庄家之前  走markerManger
func (c Tables) Start(tableId int) revel.Result {
	table := app.Tables[tableId]
	table.InitTable()
	return c.RenderJSON(table)
}

//加倍
func (c Tables) Redouble(tableId int, index int, isRedouble bool) revel.Result {
	var resultStatus = StatusRedoubleRecycle
	table := app.Tables[tableId]
	//操作次数加1
	table.MakerManager.OperateTimes++
	//底分
	if isRedouble {
		table.MakerManager.MakerIndex = index
		if table.MakerManager.Jetton == 0 {
			table.MakerManager.Jetton = app.BaseJetton
		} else {
			table.MakerManager.Jetton *= 2
		}
	}

	if table.MakerManager.OperateTimes == 4 {
		if table.MakerManager.MakerIndex == 0 {
			resultStatus = StatusRedoubleRestart
		} else {
			resultStatus = StatusRedoubleSuccess
		}
	}

	if resultStatus == StatusRedoubleSuccess {
		table.PlayTransitionManager = table.MakerManager.MakerIndex
	} else {
		table.ToNext()
	}
	result := struct {
		MarkerStatus int
		Table        *models.Table
	}{MarkerStatus: resultStatus, Table: table}

	return c.RenderJSON(result)
}

//出牌验证
func (c Tables) PlayValidation(tableId, seatIndex int, pokers []models.Poker) revel.Result {
	tablePokers := app.Tables[tableId].CurrentPokers
	tablePokersType := models.PokerArr(tablePokers.Pokers).GetType()
	currentType := models.PokerArr(pokers).GetType()

	if currentType.TypeInt == models.PokersTypeForNull {
		return c.RenderJSON(false)
	}
	if tablePokers.PlayedIndex == seatIndex || len(tablePokers.Pokers) == 0 {
		//最开始出的  或者  之前自己出的没有人打住  再出的时候没类型限制
		return c.RenderJSON(true)
	}
	//这是打别人牌  自己是炸弹
	if currentType.TypeInt == models.PokersTypeForBoom {
		if tablePokersType.TypeInt == models.PokersTypeForBoom {
			return c.RenderJSON(currentType.TypeInt == models.PokersTypeForBoom && currentType.Value > tablePokersType.Value)
		} else {
			return c.RenderJSON(true)
		}
	}

	//不是炸弹的话只能是类型相同
	if currentType.TypeInt == tablePokersType.TypeInt && len(pokers) == len(tablePokers.Pokers) {
		return c.RenderJSON(currentType.Value > tablePokersType.Value)
	}
	return c.RenderJSON(false)
}

//确定之后  出牌
func (c Tables) PlayPokers(tableId, seatIndex int, pokers []models.Poker) revel.Result {
	table := app.Tables[tableId]
	seat := table.GetSeatByIndex(seatIndex)
	seat.OutPokers = pokers
	table.HasPlayedPokers = append(table.HasPlayedPokers, pokers...)
	table.CurrentPokers = &models.CurrentPokers{
		Pokers:      pokers,
		PlayedIndex: seatIndex,
	}

	for k := 0; k < len(pokers); k++ {
		for j, po := range seat.Pokers {
			if po.Id == pokers[k].Id {
				seat.Pokers = append(seat.Pokers[:j], seat.Pokers[j+1:]...)
				break
			}
		}
	}

	//指针跳到下一个位置
	table.ToNext()

	return c.RenderJSON(table)
}
