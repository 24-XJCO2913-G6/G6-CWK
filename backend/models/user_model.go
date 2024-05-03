package models

import (
	"golang.org/x/crypto/bcrypt"
	. "main/backend/handlers"
	"time"
)

type User struct {
	Uid         int64  `xorm:"pk autoincr"`            //主键，UID
	Name        string `xorm:"NOT NULL"`               // 用户名
	Email       string `xorm:"UNIQUE NOT NULL"`        // 邮箱，唯一键
	Passwd      string `xorm:"NOT NULL"`               // 密码
	IsAdmin     bool   `xorm:"NOT NULL DEFAULT false"` // 管理员标记
	IsVIP       bool   `xorm:"NOT NULL DEFAULT false"` // VIP标记
	CreatTime   string `xorm:"UNIQUE NOT NULL"`        // 创造日期
	ProfilePhot string `xorm:""`                       // 个人头像
}

func IsExist(email string) (bool, error) {
	user := &User{Email: email}
	has, err := Db.Get(user)
	if err != nil {
		return false, err
	}
	return has, nil
}

func IsRight(email string, passwd string) (bool, error) {
	user := &User{Email: email}
	has, err := Db.Get(user)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func FindUid(email string) (int64, error) {
	user := &User{Email: email}
	has, err := Db.Get(user)
	if err != nil {
		return 0, err
	}
	if !has {
		return 0, nil
	}
	return user.Uid, nil
}

func AddUser(name string, email string, passwd string) int64 {
	user := &User{
		Email:       email,
		Name:        name,
		Passwd:      passwd,
		IsAdmin:     false,
		IsVIP:       false,
		CreatTime:   time.Now().Format("YYYY-MM-DD hh:mm:ss"),
		ProfilePhot: "",
	}

	Uid, err := Db.Insert(user)
	if err != nil {
		return Uid
	}
	return Uid
}
