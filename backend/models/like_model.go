package models

type Like struct {
	Lid  int64  `xorm:"pk autoincr"` // 轨迹的唯一标识符，被标记为主键
	Uid  int64  // 用户的唯一标识符
	Bid  int64  //blog id
	Time string //
}

func BuildModelLike() {
	err := Db.Sync2(new(Like))
	if err != nil {
		panic(err)
	}
}

func AddLike(Uid int64, Bid int64, Time string) int64 {
	like := &Like{
		Uid:  Uid,
		Bid:  Bid,
		Time: Time,
	}

	Lid, err := Db.Insert(like)
	if err != nil {
		return -1
	}
	return Lid
}
