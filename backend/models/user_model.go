package models

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	Uid     int64  `xorm:"pk autoincr"`            //主键，UID
	Name    string `xorm:"NOT NULL"`               // 用户名
	Email   string `xorm:"UNIQUE NOT NULL"`        // 邮箱，唯一键
	Passwd  string `xorm:"NOT NULL"`               // 密码
	IsAdmin bool   `xorm:"NOT NULL DEFAULT false"` // 管理员标记
}

var Slice []User

var State = make(map[string]interface{})

func IsExist(email string) bool {
	if len(Slice) == 0 {
		return false
	} else {
		//遍历切片
		for _, v := range Slice {
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
			err := bcrypt.CompareHashAndPassword([]byte(v.Passwd), []byte(passwd))
			if err != nil {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}

func FindUid(email string) int64 {
	for _, v := range Slice {
		if v.Email == email {
			return v.Uid
		}
	}
	return 0
}
