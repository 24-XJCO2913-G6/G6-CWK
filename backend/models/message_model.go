package models

type Messages struct {
	Mid         int64 `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
	Sender_id   int64
	Reveiver_id int64
	Operation   string
	Time        string
}

func BuildModelMessages() {
	err := Db.Sync2(new(Messages))
	if err != nil {
		panic(err)
	}
}

func AddMessage(Sender_id int64, Reveiver_id int64, Operation string, Time string) int64 {
	messages := &Messages{
		Sender_id:   Sender_id,
		Reveiver_id: Reveiver_id,
		Operation:   Operation,
		Time:        Time,
	}

	Mid, err := Db.Insert(messages)
	if err != nil {
		return -1
	}
	return Mid
}
