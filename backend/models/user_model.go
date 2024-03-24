package models

type User struct {
	Uid    int64  `xorm:"pk autoincr"`     //主键，UID
	Email  string `xorm:"UNIQUE NOT NULL"` // 邮箱，唯一键
	Passwd string `xorm:"NOT NULL"`        // 密码

	salt string `xorm:"VARCHAR(32)"` // 用于密码加盐哈希，注册时随机生成

	IsAdmin bool `xorm:"NOT NULL DEFAULT false"` // 管理员标记
}

var Slice []User

var State = make(map[string]interface{})

func IsExist(email string) bool {
	//如果长度为0说明尚未有用户注册
	if len(Slice) == 0 {
		return false
	} else {
		//遍历切片
		for _, v := range Slice {
			// return v.email == userEmail //此时只能和第一个比较，所以第一个之后全为false
			if v.Email == email {
				return true
			}
		}
	}
	return false
}

func IsRight(email string, passwd string) bool {
	for _, v := range Slice {
		if v.Email == email {
			//先确认姓名一致，密码相同返回true
			return v.Passwd == passwd
		}
	}
	return false
}
