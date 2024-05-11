package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

// UpdateTrack Get
func UpdateTrack(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	aToken := c.GetHeader("aToken")
	if aToken == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Track update unsuccessfully."})
	}
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)

	StrDate := c.PostForm("StrDate")
	StrTime := c.PostForm("StrTime")
	EndDate := c.PostForm("EndDate")
	EndTime := c.PostForm("EndTime")

	Duration := c.PostForm("Duration")
	Distance := c.PostForm("Distance")
	Coordinates := c.PostForm("Coordinates")

	Tid := AddTrack(Uid, StrDate, StrTime, EndDate, EndTime, Duration, Distance, Coordinates)
	clientIP := c.ClientIP()
	userAgent := c.Request.UserAgent()
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	LoId := AddLog(Uid, "UploadTrack", timeString, clientIP, userAgent)
	if Tid == -1 || LoId == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Track update unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Track update successfully."})
}

// DeliverUserTracks POST
func DeliverUserTracks(c *gin.Context) {
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		c.JSON(http.StatusOK, gin.H{"error": "Deliver Unsuccessfully"})
	}

	tracks, _ := GetTracks(Uid)

	c.JSON(http.StatusOK, gin.H{"message": "Deliver Successfully", "tracks": tracks})
}

func GetTracks(Uid int64) ([]Track, error) {
	var tracksList []Track
	err := Db.Where("uid = ?", Uid).Find(&tracksList)
	if len(tracksList) == 0 || err != nil {
		return nil, err
	}
	return tracksList, nil
}

func GetTrack(Tid int64) (Track, error) {
	var track Track
	err := Db.Where("tid = ?", Tid).Find(&track)
	if err != nil {
		return Track{}, err
	}
	return track, nil
}
