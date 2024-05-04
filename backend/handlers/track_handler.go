package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
)

// UpdateTrack Get
func UpdateTrack(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)

	StrDate := c.PostForm("#")
	StrTime := c.PostForm("#")
	EndDate := c.PostForm("#")
	EndTime := c.PostForm("#")
	Distance := c.PostForm("#")
	Coordinates := c.PostForm("#")

	Tid := AddTrack(Uid, StrDate, StrTime, EndDate, EndTime, Distance, Coordinates)
	if Tid == -1 {
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
	if err != nil {
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