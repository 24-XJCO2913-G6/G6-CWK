package models

type LikedApp struct {
	Lid      int64 `xorm:"pk autoincr"`
	Uid      int64
	UserName string
	Photo    string
	BlogName string
	Time     string
}

func BuildModelLikeApp() {
	err := Db.Sync2(new(LikedApp))
	if err != nil {
		panic(err)
	}
}

func AddLikeApp(Uid int64, UserName string, Photo string, BlogName string, Time string) int64 {
	likedApp := &LikedApp{
		Uid:      Uid,
		UserName: UserName,
		Photo:    Photo,
		BlogName: BlogName,
		Time:     Time,
	}

	Lid, err := Db.Insert(likedApp)
	if err != nil {
		return -1
	}
	return Lid
}
