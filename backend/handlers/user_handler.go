package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "main/backend/models"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	email := c.Request.FormValue("email")
	name := c.Request.FormValue("name")
	passwd := c.Request.FormValue("passwd")
	rePasswd := c.Request.FormValue("rePasswd")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name"})
		return
	}

	if !IsValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if passwd != rePasswd {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords does not match"})
		return
	}

	if IsExist(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email has been used"})
		return
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	Uid := AddUser(name, email, string(hashedPasswd))

	atoken, rtoken, _ := SetToken(strconv.FormatInt(Uid, 10), string(hashedPasswd))

	c.Set("Uid", Uid)
	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully", "aToken": atoken, "rToken": rtoken})
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

	Uid := FindUid(email)

	hashedPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	aToken, rToken, _ := SetToken(strconv.FormatInt(Uid, 10), string(hashedPasswd))

	c.Set("Uid", Uid)
	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "aToken": aToken, "rToken": rToken})
}

func AddUser(name string, email string, passwd string) int64 {
	var user User
	user.Uid = int64(len(Slice) + 1)
	user.Email = email
	user.Name = name
	user.Passwd = passwd
	user.IsAdmin = false

	Slice = append(Slice, user)
	return user.Uid
}
