package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
)

func Register(c *gin.Context) {
	email := c.Request.FormValue("Email")
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	rePasswd := c.Request.FormValue("RePasswd")

	if name == " " {
		State["state"] = 0
		State["text"] = "Invalid name!"

	} else if !IsValidEmail(email) {
		State["state"] = 0
		State["text"] = "Invalid email!"

	} else if passwd != rePasswd {
		State["state"] = 0
		State["text"] = "Wrong password repetition!"
	} else {
		Flag := IsExist(email)
		if Flag {
			State["state"] = 0
			State["text"] = "Email had been used!"
		} else {
			//用户不存在即添加用户
			AddUser(name, email, passwd)
			State["state"] = 1
			State["text"] = "Register Successfully!"
		}
	}
	c.XML(http.StatusOK, State)
}

func Login(c *gin.Context) {
	email := c.Request.FormValue("Email")
	passwd := c.Request.FormValue("Passwd")
	//先判断用户是否存在，存在再判断密码是否正确
	Bool := IsExist(email)
	if Bool {
		Bool_Pwd := IsRight(email, passwd)
		if Bool_Pwd {
			State["state"] = 1
			State["text"] = "Login Successfully!"
		} else {
			State["state"] = 0
			State["text"] = "Wrong Password!"
		}
	} else {
		State["state"] = 0
		State["text"] = "Unregister Email!"
	}

	c.XML(http.StatusOK, State)
}

func AddUser(name string, email string, passwd string) {
	var user User
	user.Uid = int64(len(Slice) + 1)
	user.Email = email
	user.Name = name
	user.Passwd = passwd
	user.IsAdmin = false
	Slice = append(Slice, user)
}
