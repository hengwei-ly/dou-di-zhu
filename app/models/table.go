package models

import (
	"log"
	"math/rand"
	"time"
)

type Table struct {
	Id                    int
	Seats                 []*Seat
	MakerManager          *MakerManager
	CurrentPokers         *CurrentPokers
	HasPlayedPokers       []Poker
	PlayTransitionManager int
	roomId                int
	baseJetton            int
}

type CurrentPokers struct {
	Pokers      []Poker
	PlayedIndex int
}

//初始化
func (t *Table) InitTable() {
	randInt := rand.Intn(2) + 1
	t.PlayTransitionManager = randInt%len(t.Seats) + 1
	t.MakerManager = &MakerManager{Jetton: t.baseJetton}
	t.CurrentPokers = &CurrentPokers{}
	t.HasPlayedPokers = []Poker{}
}

func (t *Table) AddPlayedPokers(pokers []Poker) {
	t.HasPlayedPokers = append(t.HasPlayedPokers, pokers...)
}

func (t *Table) GetSeatByIndex(seatIndex int) *Seat {
	for k := 0; k < len(t.Seats); k++ {
		if t.Seats[k].Index == seatIndex {
			return t.Seats[k]
		}
	}
	log.Println(time.Now().Format("2006-01-02 15:04:05"), " 获取座位信息失败:seatIndex = ", seatIndex)
	return nil
}

func (t *Table) ToNext() {
	t.PlayTransitionManager = (t.PlayTransitionManager+1)%3 + 1
}

func InitTables(total, jetton, roomId int) map[int]*Table {
	if total == 0 {
		total = 1
	}
	tables := map[int]*Table{}
	for i := 1; i <= total; i++ {
		tables[i] = &Table{
			Id:         i,
			Seats:      []*Seat{{Index: 1}, {Index: 2}, {Index: 3}},
			baseJetton: jetton,
			roomId:     roomId,
		}
	}
	return tables
}
