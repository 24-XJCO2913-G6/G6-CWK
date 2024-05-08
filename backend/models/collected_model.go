package models

type CollectedApp struct {
	Cid      int64 `xorm:"pk"`
	Uid      int64
	UserName int64
	Photo    string
	BlogName string
	Time     string
}

func BuildModelCollectApp() {
	err := Db.Sync2(new(CollectedApp))
	if err != nil {
		panic(err)
	}
}

func AddCollectApp(Uid int64, UserName int64, Photo string, BlogName string, Time string) int64 {
	collectedApp := &CollectedApp{
		Uid:      Uid,
		UserName: UserName,
		Photo:    Photo,
		BlogName: BlogName,
		Time:     Time,
	}

	Cid, err := Db.Insert(collectedApp)
	if err != nil {
		return -1
	}
	return Cid
}
