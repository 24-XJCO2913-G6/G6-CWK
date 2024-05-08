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
	err := c.Request.ParseForm()
	if err != nil {
		return
	}
	Claim, _ := CheckToken(c.GetString("aToken"))
	Uid_tmp := Claim.Uid
	visibility_tmp := c.PostForm("visibility")
	currentTime := time.Now()
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
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
	//TODO visibility 没传
	c.JSON(http.StatusOK, gin.H{"message": "Blog publish Successfully"})
}
