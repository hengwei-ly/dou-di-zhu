package models

import (
	"sort"
)

//返回每个人的牌和底牌
func GetDoudizhuPokers(pers []*Seat) []Poker {
	initPokers := GetDisorderPokers()
	nums := (len(initPokers) - 3) / len(pers)

	for k := 0; k < len(pers); k++ {
		pokers := initPokers[k*nums : (k+1)*nums]
		sort.Sort(PokerArr(pokers))
		pers[k].Pokers = pokers
	}

	return initPokers[len(initPokers)-3:]
}
