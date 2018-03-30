package models

import (
	"math/rand"
	"time"
)

type Poker struct {
	Id    int    `json:"id" xorm:"value"`
	Value int    `json:"value" xorm:"value"`
	HuaSe int    `json:"hua_se" xorm:"hua_se"`
	Skin  string `json:"skin" xorm:"skin"`
}

//qq  这里适合用工厂模式 暂时先写死
func GetDisorderPokers() []Poker {
	sortedPokers := initPokers()
	pokers := make([]Poker, len(sortedPokers))

	ra := rand.New(rand.NewSource(time.Now().Unix()))
	for k := 0; k < len(pokers); k++ {
		for {
			idx := ra.Intn(len(sortedPokers))
			pokers[k] = sortedPokers[idx]
			sortedPokers = append(sortedPokers[:idx], sortedPokers[idx+1:]...)
			break
		}
	}
	return pokers
}

//给几副牌
func initPokers() []Poker {
	pokers := make([]Poker, 54)
	for i := 1; i <= 13; i++ {
		for j := 1; j <= 4; j++ {
			index := (i-1)*4 + j - 1
			pokers[index] = Poker{
				Id:    index + 1,
				Value: i,
				HuaSe: j,
			}
		}
	}
	pokers[52] = Poker{Id: 53, Value: 14, HuaSe: 0, Skin: ""}
	pokers[53] = Poker{Id: 54, Value: 15, HuaSe: 0, Skin: ""}
	return pokers
}

type PokersType struct {
	TypeInt PokerTypeInt
	Value   int
}

func (pt PokersType) IsMaxInCurrentTypes() bool {
	switch pt.TypeInt {
	case PokersTypeForBoom:
		return pt.Value == 14
	case PokersTypeForSingle:
		return pt.Value == 15
	case PokersTypeForDouble, PokersTypeForTrebleWithNull, PokersTypeForTrebleWithSingle, PokersTypeForTrebleWithDouble,
		PokersTypeForQuadrupleWithSingle, PokersTypeForQuadrupleWithDouble, PokersTypeForQuadrupleWithDouble2:
		return pt.Value == 13
	default:
		return pt.Value == 12
	}
}

type PokerTypeInt int

func (t PokerTypeInt) String() string {
	switch t {
	case PokersTypeForNull:
		return "没有类型"
	case PokersTypeForSingle:
		return "单张"
	case PokersTypeForDouble:
		return "对子"
	case PokersTypeForTrebleWithNull:
		return "三不带"
	case PokersTypeForTrebleWithSingle:
		return "三带一"
	case PokersTypeForTrebleWithDouble:
		return "三带二"
	case PokersTypeForQuadrupleWithSingle:
		return "四带二"
	case PokersTypeForQuadrupleWithDouble:
		return "四带一对"
	case PokersTypeForQuadrupleWithDouble2:
		return "四带两对"
	case PokersTypeForSerialSingle:
		return "顺子"
	case PokersTypeForSerialDouble:
		return "连对"
	case PokersTypeForSerialTrebleWithNull:
		return "飞机"
	case PokersTypeForSerialTrebleWithSingle:
		return "飞机带单张"
	case PokersTypeForSerialTrebleWithDouble:
		return "飞机带对子"
	case PokersTypeForBoom:
		return "炸弹"
	}
	return "未知类型"
}

const (
	PokersTypeForNull = iota
	PokersTypeForDouble
	PokersTypeForTrebleWithNull
	PokersTypeForTrebleWithSingle
	PokersTypeForSingle
	PokersTypeForTrebleWithDouble
	PokersTypeForQuadrupleWithSingle
	PokersTypeForQuadrupleWithDouble
	PokersTypeForQuadrupleWithDouble2
	PokersTypeForSerialSingle
	PokersTypeForSerialDouble
	PokersTypeForSerialTrebleWithNull
	PokersTypeForSerialTrebleWithSingle
	PokersTypeForSerialTrebleWithDouble
	PokersTypeForBoom
)
