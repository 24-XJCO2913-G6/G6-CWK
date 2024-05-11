package models

import "time"

type Order struct {
	Oid     int64 `xorm:"pk autoincr"`
	Uid     int64
	Content string
	Time    time.Time
	Price   string
	State   string
}

func BuildModelOrder() {
	err := Db.Sync2(new(Order))
	if err != nil {
		panic(err)
	}
}

func AddOrder(Uid int64, Content string, Time time.Time, Price string, State string) int64 {
	order := &Order{
		Uid:     Uid,
		Content: Content,
		Time:    Time,
		Price:   Price,
		State:   State,
	}

	Oid, err := Db.Insert(order)
	if err != nil {
		return -1
	}
	return Oid
}
