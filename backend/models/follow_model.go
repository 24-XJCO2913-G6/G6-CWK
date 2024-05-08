package models

type Follow struct {
	Fid        int64 `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
	FollowId   int64
	FollowById int64
	Time       string
}

func BuildModelFollow() {
	err := Db.Sync2(new(Follow))
	if err != nil {
		panic(err)
	}
}

func AddFollow(FollowId int64, FollowById int64, Time string) int64 {
	follow := &Follow{
		FollowId:   FollowId,
		FollowById: FollowById,
		Time:       Time,
	}

	Fid, err := Db.Insert(follow)
	if err != nil {
		return -1
	}
	return Fid
}
