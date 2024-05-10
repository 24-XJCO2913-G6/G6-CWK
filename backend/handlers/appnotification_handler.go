package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"strconv"
)

func GetLikeMes(c *gin.Context) ([]LikedApp, error) {
	var likedMessages []LikedApp
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	err := Db.Where("uid = ?", Uid).Find(&likedMessages)
	if err != nil {
		return []LikedApp{}, err
	}
	// 输出结果

	return likedMessages, nil
}

func GetCollectMes(c *gin.Context) ([]CollectedApp, error) {
	var collectMessages []CollectedApp
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	err := Db.Where("uid = ?", Uid).Find(&collectMessages)
	if err != nil {
		return []CollectedApp{}, err
	}
	// 输出结果

	return collectMessages, nil
}

func GetReviewMes(c *gin.Context) ([]ReviewedApp, error) {
	var reviewMessages []ReviewedApp
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	err := Db.Where("uid = ?", Uid).Find(&reviewMessages)
	if err != nil {
		return []ReviewedApp{}, err
	}
	// 输出结果

	return reviewMessages, nil
}
