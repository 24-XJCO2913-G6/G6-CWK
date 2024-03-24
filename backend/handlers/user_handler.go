package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
)

func Register(c *gin.Context) {
	email := c.Request.FormValue("Email")
	passwd := c.Request.FormValue("Passwd")

	Flag := IsExist(email)
	if Flag {
		State["state"] = 1
		State["text"] = "Email had been used!"
	} else {
		//用户不存在即添加用户
		AddUser(email, passwd)
		State["state"] = 1
		State["text"] = "Register Successfully!"
	}
	c.String(http.StatusOK, "%v", State)
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
		State["state"] = 2
		State["text"] = "Unregister Email!"
	}

	c.String(http.StatusOK, "%v", State)
}

func AddUser(email string, passwd string) {
	var user User
	user.Email = email
	user.Passwd = passwd
	user.Uid = int64(len(Slice) + 1)
	user.IsAdmin = false
	Slice = append(Slice, user)
}
