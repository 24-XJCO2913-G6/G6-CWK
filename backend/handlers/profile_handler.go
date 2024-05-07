package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"strconv"
)

func GetName(c *gin.Context) (string, error) {

	var users []User
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&users)
		if err != nil {
			return "null", err
		}
	}
	Name := users[0].Name
	return Name, nil
}

func GetEmail(c *gin.Context) (string, error) {

	var users []User
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&users)
		if err != nil {
			return "null", err
		}
	}
	Email := users[0].Email
	return Email, nil
}

func GetPhoto(c *gin.Context) (string, error) {

	var users []User
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&users)
		if err != nil {
			return "null", err
		}
	}
	ProfilePhot := users[0].ProfilePhot
	return ProfilePhot, nil
}
func GetSignature(c *gin.Context) (string, error) {

	var users []User
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&users)
		if err != nil {
			return "null", err
		}
	}
	Signature := users[0].Signature
	return Signature, nil
}
