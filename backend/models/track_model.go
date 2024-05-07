package models

type Track struct {
	Tid         int64 `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
	Uid         int64 // 用户的唯一标识符
	StrDate     string
	StrTime     string // 轨迹的起始时间
	EndDate     string
	EndTime     string // 轨迹的结束时间
	Distance    string // 轨迹的距离
	Coordinates string // 轨迹的坐标信息
	// 如[x,y],[x,y]
}
type Records struct {
	Uid      int64
	Distance string
	Name     string
	Photo    string
}

func BuildModelTrack() {
	err := Db.Sync2(new(Track))
	if err != nil {
		panic(err)
	}
}

func AddTrack(Uid int64, StrDate string, StrTime string, EndDate string, EndTime string, Distance string, Coordinates string) int64 {
	track := &Track{
		Uid:         Uid,
		StrDate:     StrDate,
		StrTime:     StrTime,
		EndDate:     EndDate,
		EndTime:     EndTime,
		Distance:    Distance,
		Coordinates: Coordinates,
	}

	Tid, err := Db.Insert(track)
	if err != nil {
		return -1
	}
	return Tid
}
