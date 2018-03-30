package models

import (
	"doudizhu/app/utils"
	"sort"
)

//多张牌
type PokerArr []Poker

func (p PokerArr) Len() int {
	return len(p)
}
func (p PokerArr) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].HuaSe > p[j].HuaSe
	}
	return p[i].Value > p[j].Value
}
func (p PokerArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PokerArr) Reverse() {
	for k := 0; k < len(p)/2; k++ {
		p[k], p[len(p)-1-k] = p[len(p)-1-k], p[k]
	}
}

//按照显示顺序排序
func (p PokerArr) SortForPlay() {
	if len(p) > 3 {
		repeats := p.GetRepeats()
		sort.Slice(p, func(i, j int) bool {
			if len(repeats) > 0 {
				if repeats[p[i].Value] != repeats[p[j].Value] {
					return repeats[p[i].Value] > repeats[p[j].Value]
				}
			}
			if p[i].Value != p[j].Value {
				return p[i].Value < p[j].Value
			}
			return p[i].HuaSe < p[j].HuaSe
		})
	} else {
		sort.Sort(p)
		p.Reverse()
	}
}

//找出数组里面重复的值的个数  返回 value:num
func (p PokerArr) GetRepeats() map[int]int {
	repeats := map[int]int{}
	for i := 0; i < len(p); i++ {
		val, ok := repeats[p[i].Value]
		if ok {
			repeats[p[i].Value] = val + 1
		} else {
			repeats[p[i].Value] = 1
		}
	}
	return repeats
}

/*
	获取单位长度的首个对象的集合  看是不是连续  顺子 unitLen = 1 连对:2 飞机:3  飞机带一张:4  飞机带两张:5*
*/
func (p PokerArr) GetSerialMaxValue(unitLength int) (bool, int) {
	arr := make([]int, len(p)/unitLength)
	for i := 0; i < len(p); i = i + unitLength {
		arr[i/unitLength] = p[i].Value
	}
	max := utils.IsSerial(arr)
	return max > 0 && max < 14, max
}

func (p PokerArr) IsTrebleWithDouble() (PokerTypeInt, int) {
	if len(p) != 5 {
		return PokersTypeForNull, 0
	}
	if p[0].Value == p[1].Value && p[0].Value == p[2].Value &&
		p[0].Value != p[3].Value && p[3].Value == p[4].Value {
		return PokersTypeForTrebleWithDouble, p[0].Value
	}
	return PokersTypeForNull, 0
}

//判断是不是顺子  不能到2
func (p PokerArr) IsSerialSingle() (PokerTypeInt, int) {
	if len(p) < 5 {
		return PokersTypeForNull, 0
	}

	if ok, val := p.GetSerialMaxValue(1); ok {
		return PokersTypeForSerialSingle, val
	}
	return PokersTypeForNull, 0
}

//按照play方式排序之后的数组
func (p PokerArr) IsSerialDouble() (PokerTypeInt, int) {
	if len(p) < 6 || len(p)%2 != 0 {
		return PokersTypeForNull, 0
	}

	if ok, val := p.GetSerialMaxValue(2); ok {
		return PokersTypeForSerialDouble, val
	}
	return PokersTypeForNull, 0
}

//按照play方式排序之后的数组
func (p PokerArr) IsSerialTrebleWithNull() (PokerTypeInt, int) {
	if len(p) < 6 || len(p)%3 != 0 {
		return PokersTypeForNull, 0
	}
	if ok, val := p.GetSerialMaxValue(3); ok {
		return PokersTypeForSerialTrebleWithNull, val
	}
	return PokersTypeForNull, 0
}

//按照play方式排序之后的数组
func (p PokerArr) IsSerialTrebleWithSingle() (PokerTypeInt, int) {
	if len(p) < 8 || len(p)%4 != 0 {
		return PokersTypeForNull, 0
	}
	repeats := p.GetRepeats()
	if len(repeats)%2 != 0 {
		return PokersTypeForNull, 0
	}
	threeNums, withNums := 0, 0
	for _, num := range repeats {
		if num == 3 {
			threeNums++
		} else if num == 1 {
			withNums++
		} else {
			return PokersTypeForNull, 0
		}
	}
	if threeNums != withNums {
		return PokersTypeForNull, 0
	}
	if ok, val := p.GetSerialMaxValue(4); ok {
		return PokersTypeForSerialTrebleWithSingle, val
	}
	return PokersTypeForNull, 0
}

func (p PokerArr) IsSerialTrebleWithDouble() (PokerTypeInt, int) {
	if len(p) < 8 || len(p)%5 != 0 {
		return PokersTypeForNull, 0
	}
	repeats := p.GetRepeats()
	if len(repeats)%2 != 0 {
		return PokersTypeForNull, 0
	}

	threeNums, withNums := 0, 0
	for _, num := range repeats {
		if num == 3 {
			threeNums++
		} else if num == 2 {
			withNums++
		} else {
			return PokersTypeForNull, 0
		}
	}

	if threeNums != withNums {
		return PokersTypeForNull, 0
	}
	if ok, val := p.GetSerialMaxValue(5); ok {
		return PokersTypeForSerialTrebleWithDouble, val
	}
	return PokersTypeForNull, 0
}

func (p PokerArr) IsQuadrupleWithSingle() (PokerTypeInt, int) {
	if len(p) != 6 {
		return PokersTypeForNull, 0
	}
	if p[0].Value == p[1].Value && p[0].Value == p[2].Value &&
		p[0].Value == p[3].Value && p[4].Value != p[5].Value {
		return PokersTypeForQuadrupleWithSingle, p[0].Value
	}
	return PokersTypeForNull, 0
}

func (p PokerArr) IsQuadrupleWithDouble() (PokerTypeInt, int) {
	if len(p) != 6 {
		return PokersTypeForNull, 0
	}
	if p[0].Value == p[1].Value && p[0].Value == p[2].Value &&
		p[0].Value == p[3].Value && p[4].Value == p[5].Value {
		return PokersTypeForQuadrupleWithDouble, p[0].Value
	}
	return PokersTypeForNull, 0
}

func (p PokerArr) IsQuadrupleWithDouble2() (PokerTypeInt, int) {
	if len(p) != 8 {
		return PokersTypeForNull, 0
	}
	if p[0].Value == p[1].Value && p[0].Value == p[2].Value &&
		p[0].Value == p[3].Value && p[4].Value == p[5].Value &&
		p[5].Value != p[6].Value && p[6].Value == p[7].Value {
		return PokersTypeForQuadrupleWithDouble2, p[0].Value
	}
	return PokersTypeForNull, 0
}

//获取类型
func (ps PokerArr) GetType() (pt PokersType) {
	if !sort.IsSorted(ps) {
		sort.Sort(ps)
	}

	switch len(ps) {
	case 1:
		pt.Value = ps[0].Value
		pt.TypeInt = PokersTypeForSingle
		return
	case 2:
		//炸
		if ps[0].HuaSe == 0 && ps[1].HuaSe == 0 {
			pt.Value = 16
			pt.TypeInt = PokersTypeForBoom
		} else if ps[0].Value == ps[1].Value {
			pt.Value = ps[0].Value
			pt.TypeInt = PokersTypeForDouble
		}
		return
	case 3:
		if ps[0].Value == ps[1].Value && ps[0].Value == ps[2].Value {
			pt.Value = ps[0].Value
			pt.TypeInt = PokersTypeForTrebleWithNull
		}
		return
	case 4:
		ps.SortForPlay()
		if ps[0].Value == ps[1].Value && ps[0].Value == ps[2].Value {
			pt.Value = ps[0].Value
			if ps[0].Value == ps[3].Value {
				pt.TypeInt = PokersTypeForBoom
			} else {
				pt.TypeInt = PokersTypeForTrebleWithSingle
			}
		}
		return
	case 5:
		ps.SortForPlay()
		for _, flag := range []func() (PokerTypeInt, int){
			ps.IsSerialSingle,
			ps.IsTrebleWithDouble,
		} {
			if typ, val := flag(); typ != PokersTypeForNull {
				pt.TypeInt = typ
				pt.Value = val
			}
		}
		return
	case 6:
		ps.SortForPlay()
		for _, flag := range []func() (PokerTypeInt, int){
			ps.IsSerialSingle,
			ps.IsSerialDouble,
			ps.IsSerialTrebleWithNull,
			ps.IsQuadrupleWithSingle,
			ps.IsQuadrupleWithDouble,
		} {
			if typ, val := flag(); typ != PokersTypeForNull {
				pt.TypeInt = typ
				pt.Value = val
			}
		}
		return
	case 8:
		ps.SortForPlay()
		for _, flag := range []func() (PokerTypeInt, int){
			ps.IsSerialSingle,
			ps.IsSerialDouble,
			ps.IsSerialTrebleWithSingle,
			ps.IsQuadrupleWithDouble2,
		} {
			if typ, val := flag(); typ != PokersTypeForNull {
				pt.TypeInt = typ
				pt.Value = val
			}
		}
		return
	default:
		//大于5 且不包含6和8 的
		ps.SortForPlay()
		for _, flag := range []func() (PokerTypeInt, int){
			ps.IsSerialSingle,
			ps.IsSerialDouble,
			ps.IsSerialTrebleWithNull,
			ps.IsSerialTrebleWithSingle,
			ps.IsSerialTrebleWithDouble,
		} {
			if typ, val := flag(); typ != PokersTypeForNull {
				pt.TypeInt = typ
				pt.Value = val
			}
		}
	}
	return
}
