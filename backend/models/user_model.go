package models

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	Uid    int64  `xorm:"pk autoincr"`     //主键，UID
	Name   string `xorm:"NOT NULL"`        // 用户名
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
	// 正则表达式验证邮箱格式
	// 标准的电子邮件地址格式包含一个或多个字母、数字、下划线、连字符和点，后跟 @ 符号，
	// 然后是一个或多个字母、数字、连字符和点，最后是一个点和两个到四个字母
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}
