package models

type Vip struct {
	Vid     int64 `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
	Uid     int64
	StrTime string
	EndTime string
}

func BuildModelVip() {
	err := Db.Sync2(new(Vip))
	if err != nil {
		panic(err)
	}
}

func AddVip(Uid int64, StrTime string, EndTime string) int64 {
	vip := &Vip{
		Uid:     Uid,
		StrTime: StrTime,
		EndTime: EndTime,
	}

	Vid, err := Db.Insert(vip)
	if err != nil {
		return -1
	}
	return Vid
}
