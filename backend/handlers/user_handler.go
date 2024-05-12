package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "main/backend/models"
	"net/http"
	"regexp"
	"strconv"
)

func Register(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	email := c.PostForm("email")
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	rePasswd := c.PostForm("rePasswd")

	if name == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid name"})
		return
	}

	if !IsValidEmail(email) {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid email"})
		return
	}

	if passwd != rePasswd {
		c.JSON(http.StatusOK, gin.H{"error": "Passwords does not match"})
		return
	}

	isExit, err := IsExist(email)
	if isExit {
		c.JSON(http.StatusOK, gin.H{"error": "Email has been used"})
		return
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to hash password"})
		return
	}
	AddUser(name, email, string(hashedPasswd))

	c.JSON(http.StatusOK, gin.H{"message": "Register Successfully"})
}

func Login(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	email := c.PostForm("email")
	//fmt.Printf("type: %T\n", email)
	passwd := c.PostForm("passwd")

	isEmail, err := IsExist(email)
	if !isEmail {
		c.JSON(http.StatusOK, gin.H{"error": "Unregistered email"})
		return
	}

	isRight, err := IsRight(email, passwd)
	if !isRight {
		c.JSON(http.StatusOK, gin.H{"error": "Wrong password"})
		return
	}

	Uid, _ := FindUid(email)

	hashedPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	aToken, rToken, _ := SetToken(strconv.FormatInt(Uid, 10), string(hashedPasswd))

	c.Set("Uid", Uid)
	c.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "aToken": aToken, "rToken": rToken, "Uid": Uid})
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}
