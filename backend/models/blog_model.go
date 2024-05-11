package models

type Blog struct {
	Id         int64 `xorm:"pk autoincr"` // 轨迹的唯一标识符，被标记为主键
	Uid        int64 // 用户的唯一标识符
	Pub_time   string
	Visibility int64 //
	Content    string
	Picture    string //
	Title      string //
	Tid        int64
}

type Blog_display struct {
	Bid          int64
	Author       string
	Photo        string
	Pub_time     string
	LikeCount    int64
	CollectCount int64
	CommentCount int64
	Visibility   int64 //
	Content      string
	Picture      string //
	Title        string //
	IsFollow     int64
	Coordinates  string
	AuthorId     int64
}

func BuildModelBlog() {
	err := Db.Sync2(new(Blog))
	if err != nil {
		panic(err)
	}
}

func AddBlog(Uid int64, Pub_time string, Visibility int64, Content string, Picture string, Title string, Tid int64) int64 {
	blog := &Blog{
		Uid:        Uid,
		Pub_time:   Pub_time,
		Visibility: Visibility,
		Content:    Content,
		Picture:    Picture,
		Title:      Title,
		Tid:        Tid,
	}

	Id, err := Db.Insert(blog)
	if err != nil {
		return -1
	}
	return Id
}
