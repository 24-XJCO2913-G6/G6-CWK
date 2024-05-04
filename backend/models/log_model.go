package models

type Log struct {
	Lid  int64 `xorm:"pk"`
	Uid  int64
	Ope  string
	Time string
}

func BuildModelLog() {
	err := Db.Sync2(new(Log))
	if err != nil {
		panic(err)
	}
}

func AddLog(Uid int64, Ope string, Time string) int64 {
	log := &Log{
		Uid:  Uid,
		Ope:  Ope,
		Time: Time,
	}

	Lid, err := Db.Insert(log)
	if err != nil {
		return -1
	}
	return Lid
}
