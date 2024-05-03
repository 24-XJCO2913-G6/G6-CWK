package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Uid         string `xorm:"pk"`                     //主键，UID
	Name        string `xorm:"NOT NULL"`               // 用户名
	Email       string `xorm:"UNIQUE NOT NULL"`        // 邮箱，唯一键
	Passwd      string `xorm:"NOT NULL"`               // 密码
	IsAdmin     bool   `xorm:"NOT NULL DEFAULT false"` // 管理员标记
	IsVip       bool   `xorm:"NOT NULL DEFAULT false"` // VIP标记
	CreatedTime string `xorm:"UNIQUE NOT NULL"`        // 创造日期
	ProfilePhot string `xorm:""`                       // 个人头像
}

func BuildModelUser() {
	err := Db.Sync2(new(User))
	if err != nil {
		panic(err)
	}
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

func FindUid(email string) (string, error) {
	user := &User{Email: email}
	has, err := Db.Get(user)
	if err != nil {
		return "-1", err
	}
	if !has {
		return "-1", nil
	}
	return user.Uid, nil
}

func AddUser(name string, email string, passwd string) int64 {
	user := &User{
		Email:       email,
		Name:        name,
		Passwd:      passwd,
		IsAdmin:     false,
		IsVip:       false,
		CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
		ProfilePhot: "",
	}

	Uid, err := Db.Insert(user)
	if err != nil {
		return Uid
	}
	return Uid
}
