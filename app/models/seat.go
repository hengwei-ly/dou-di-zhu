package models

type Seat struct {
	Index     int
	Pokers    []Poker
	OutPokers []Poker
	IsAuto    bool
	UserId    int
}
