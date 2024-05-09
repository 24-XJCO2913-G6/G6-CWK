package models

type Vip struct {
	Vid     int64 `xorm:"pk autoincr"` // 轨迹的唯一标识符，被标记为主键
	Uid     int64
	StrTime string
	EndTime string
}

func BuildModelVip() {
	err := Db.Sync2(new(Vip))
	if err != nil {
		panic(err)
	}
}

func AddVip(Uid int64, StrTime string, EndTime string) int64 {
	vip := &Vip{
		Uid:     Uid,
		StrTime: StrTime,
		EndTime: EndTime,
	}

	Vid, err := Db.Insert(vip)
	if err != nil {
		return -1
	}
	return Vid
}

func GetVip(Uid int64) (Vip, error) {
	var vip Vip
	err := Db.Where("Uid = ?", Uid).Find(&vip)
	if err != nil {
		return Vip{}, err
	}
	return vip, nil
}

func IsVip(Uid int64) (int64, error) {
	var vip Vip
	has, err := Db.Where("Uid = ?", Uid).Get(&vip)
	if err != nil {
		return 0, err
	} else if !has {
		return 0, nil
	} else {
		return 1, nil
	}
}
