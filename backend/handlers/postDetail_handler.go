package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

func GetBlog(Id int64) (Blog, error) {
	var blog Blog
	err := Db.Where("Id = ?", Id).Find(&blog)
	if err != nil {
		return Blog{}, err
	}
	return blog, nil
}

func GetReviews(Id int64) ([]ReviewInBlog, error) {
	var reviews []Review
	var reviews_in_blog []ReviewInBlog

	err := Db.Where("Bid = ?", Id).Find(&reviews)
	if err != nil {
		return []ReviewInBlog{}, err
	}
	for _, review := range reviews {
		var reviewers []User
		err := Db.Where("Uid = ?", review.Uid).Find(&reviewers)
		if err != nil {
			return []ReviewInBlog{}, err
		}
		reviewInBlog := ReviewInBlog{Time: review.Time, Content: review.Content, Reviewer: reviewers[0].Name, Reviewer_photo: reviewers[0].ProfilePhot}
		reviews_in_blog = append(reviews_in_blog, reviewInBlog)
	}
	return reviews_in_blog, nil
}

var track_tmp []Track

func BlogPublish(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	visibility_tmp := c.PostForm("visibility")
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	title := c.PostForm("title")
	description := c.PostForm("content")
	photos := c.PostForm("imgs")
	Tid_tmp := c.PostForm("route_id")
	Tid, err := strconv.ParseInt(Tid_tmp, 10, 64)
	if err != nil {
		fmt.Println("Tid transfer fail:", err)
		return
	}
	visibility, err := strconv.ParseInt(visibility_tmp, 10, 64)
	if err != nil {
		fmt.Println("Visibility transfer fail:", err)
		return
	}
	err2 := Db.Where("Tid = ?", Tid).Find(&track_tmp)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid Tid, not exists"})
		return
	}
	if title == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Need title"})
		return
	}
	if description == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Need description"})
		return
	}
	AddBlog(Uid, timeString, visibility, description, photos, title, Tid)
	c.JSON(http.StatusOK, gin.H{"message": "Blog publish Successfully"})
}

func GetFriendsBlogs(uid int64) ([]Blog, error) {
	var friends []Friend
	var allBlogs []Blog

	// 获取指定用户的所有朋友
	friends, err := GetFriends(uid)
	if err != nil {
		return nil, err
	}

	// 遍历每个朋友，查询其发布的所有博客
	for _, friend := range friends {
		var blogs []Blog
		err := Db.Where("uid = ?", friend.Uid).Find(&blogs)
		if err != nil {
			return nil, err
		}
		allBlogs = append(allBlogs, blogs...)
	}

	return allBlogs, nil
}

func GetLikedMess(Uid int64) ([]LikedApp, error) {

	var likeMess []LikedApp
	err := Db.Where("uid = ?", Uid).Find(&likeMess)
	if err != nil {
		return nil, err
	}

	return likeMess, nil
}
func GetCollectedMess(Uid int64) ([]CollectedApp, error) {

	var collectMess []CollectedApp
	err := Db.Where("uid = ?", Uid).Find(&collectMess)
	if err != nil {
		return nil, err
	}

	return collectMess, nil
}
func GetReviewedMess(Uid int64) ([]ReviewedApp, error) {

	var reviewMess []ReviewedApp
	err := Db.Where("uid = ?", Uid).Find(&reviewMess)
	if err != nil {
		return nil, err
	}

	return reviewMess, nil
}
