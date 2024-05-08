package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Uid         int64  `xorm:"pk"`                     //主键，UID
	Name        string `xorm:"NOT NULL"`               // 用户名
	Email       string `xorm:"UNIQUE NOT NULL"`        // 邮箱，唯一键
	Passwd      string `xorm:"NOT NULL"`               // 密码
	IsAdmin     bool   `xorm:"NOT NULL DEFAULT false"` // 管理员标记
	IsVip       bool   `xorm:"NOT NULL DEFAULT false"` // VIP标记
	CreatedTime string `xorm:"UNIQUE NOT NULL"`        // 创造日期
	ProfilePhot string `xorm:""`                       // 个人头像
	Signature   string `xorm:""`                       // 签名
}

type Liked struct {
	Bid         int64
	Content     string
	LikeCount   int64
	ReviewCount int64
	Review      []Review
	Picture     string
	Title       string
	Visibility  int64
	Time        string
}
type Collected struct {
	Bid          int64
	Content      string
	CollectCount int64
	ReviewCount  int64
	Review       []Review
	Picture      string
	Title        string
	Visibility   int64
	Time         string
}

func BuildModelUser() {
	err := Db.Sync2(new(User))
	if err != nil {
		panic(err)
	}
}

func IsExist(email string) (bool, error) {
	user := &User{Email: email}
	has, _ := Db.Get(user)
	//if err != nil {
	//	return false, err
	//}
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
		return -1, err
	}
	if !has {
		return -1, nil
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
