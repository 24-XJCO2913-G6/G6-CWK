package models

type Collect struct {
	Cid  int64  `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
	Uid  int64  // 用户的唯一标识符
	Bid  int64  //blog id
	Time string //
}

func BuildModelCollect() {
	err := Db.Sync2(new(Collect))
	if err != nil {
		panic(err)
	}
}

func AddCollect(Uid int64, Bid int64, Time string) int64 {
	collect := &Collect{
		Uid:  Uid,
		Bid:  Bid,
		Time: Time,
	}

	Cid, err := Db.Insert(collect)
	if err != nil {
		return -1
	}
	return Cid
}
