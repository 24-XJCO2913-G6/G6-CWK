package models

type Review struct {
	Rid     int64 `xorm:"pk"`
	Uid     int64
	Bid     int64
	Time    string
	Content string
}

func BuildModelReview() {
	err := Db.Sync2(new(Review))
	if err != nil {
		panic(err)
	}
}

func AddReview(Uid int64, Bid int64, Time string, Content string) int64 {
	review := &Review{
		Uid:     Uid,
		Bid:     Bid,
		Time:    Time,
		Content: Content,
	}

	Rid, err := Db.Insert(review)
	if err != nil {
		return -1
	}
	return Rid
}
