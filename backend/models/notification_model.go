package models

type Notification struct {
	Nid     int64 `xorm:"pk autoincr"`
	Uid     int64
	Sender  int64
	Ope     string
	Time    string
	Content string
	Photo   string
}

func BuildModelNotification() {
	err := Db.Sync2(new(Notification))
	if err != nil {
		panic(err)
	}
}

func AddNotification(Uid int64, Sender int64, Ope string, Time string, Content string, Photo string) int64 {
	notification := &Notification{
		Uid:     Uid,
		Ope:     Ope,
		Sender:  Sender,
		Time:    Time,
		Content: Content,
		Photo:   Photo,
	}

	Nid, err := Db.Insert(notification)
	if err != nil {
		return -1
	}
	return Nid
}
