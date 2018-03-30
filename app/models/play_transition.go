package models

//chupai dong zuo guanli
type PlayTransitionManager struct {
	CurrentIndex int
	NextIndex    int
}

func NewDoudizhuTranManager(marker *MakerManager) *PlayTransitionManager {
	return &PlayTransitionManager{
		CurrentIndex: marker.MakerIndex,
		NextIndex:    (marker.MakerIndex + 1) % 3,
	}
}

////返回结果 用于显示
//func (t *PlayTransitionManager) Play(seat *Seat) string {
//	if len(seat.OutPokers) == 0 {
//		return "Pass"
//	} else {
//		t.CurrentIndex = seat.Index
//		t.NextIndex = (seat.Index + 1) % 3
//		t.Table.CurrentPokers = &CurrentPokers{
//			Pokers:      seat.OutPokers,
//			PlayedIndex: seat.Index,
//		}
//		t.Table.AddPlayedPokers(seat.OutPokers)
//	}
//	return "ok"
//}
