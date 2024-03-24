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

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name"})
		return
	}

	if !IsValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if passwd != rePasswd {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	if IsExist(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email has been used"})
		return
	}

	AddUser(name, email, passwd)

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})
}

func Login(c *gin.Context) {
	email := c.Request.FormValue("Email")
	passwd := c.Request.FormValue("Passwd")

	if !IsExist(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unregistered email"})
		return
	}

	if !IsRight(email, passwd) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully"})
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
