package models

type ReviewedApp struct {
	Rid      int64 `xorm:"pk"`
	Uid      int64
	UserName string
	Photo    string
	BlogName string
	Time     string
}

func BuildModelReviewApp() {
	err := Db.Sync2(new(ReviewedApp))
	if err != nil {
		panic(err)
	}
}

func AddReviewApp(Uid int64, UserName string, Photo string, BlogName string, Time string) int64 {
	reviewedApp := &ReviewedApp{
		Uid:      Uid,
		UserName: UserName,
		Photo:    Photo,
		BlogName: BlogName,
		Time:     Time,
	}

	Rid, err := Db.Insert(reviewedApp)
	if err != nil {
		return -1
	}
	return Rid
}
