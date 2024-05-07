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

func GetLike(c *gin.Context) ([]Liked, error) {
	var likedBlogs []Liked
	var likeList []Like
	var blogLiked []Blog
	var reviews []Review
	var likeCount []Like
	var reviewAmount int64 = 0
	var likeAmount int64 = 0
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&likeList)
		if err != nil {
			return []Liked{}, err
		}
	}
	for _, liked := range likeList {
		Id := liked.Bid
		err := Db.Where("Id = ?", Id).Find(&blogLiked)
		if err != nil {
			return []Liked{}, err
		}
		err2 := Db.Where("Bid = ?", Id).Find(&reviews)
		if err2 != nil {
			return []Liked{}, err
		}
		for range reviews {
			reviewAmount += 1
		}
		err3 := Db.Where("Bid = ?", Id).Find(&likeCount)
		if err3 != nil {
			return []Liked{}, err
		}
		for range likeCount {
			likeAmount += 1
		}
		likedBlog := Liked{Bid: blogLiked[0].Id, Content: blogLiked[0].Content, Picture: blogLiked[0].Picture, Tag: blogLiked[0].Tag, Visibility: blogLiked[0].Visibility, Time: blogLiked[0].Pub_time, Review: reviews, ReviewCount: reviewAmount, LikeCount: likeAmount}
		likedBlogs = append(likedBlogs, likedBlog)
	}
	return likedBlogs, nil
}
