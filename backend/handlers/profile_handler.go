package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"strconv"
)

func GetName(Uid int64) (string, error) {

	var users []User
	if Uid != -1 {
		err1 := Db.Where("Uid = ?", Uid).Find(&users)
		if err1 != nil {
			return "null", err1
		}
	}
	Name := users[0].Name
	return Name, nil
}

func GetEmail(Uid int64) (string, error) {

	var users []User
	if Uid != -1 {
		err1 := Db.Where("Uid = ?", Uid).Find(&users)
		if err1 != nil {
			return "null", err1
		}
	}
	Email := users[0].Email
	return Email, nil
}

func GetPhoto(Uid int64) (string, error) {

	var users []User
	if Uid != -1 {
		err1 := Db.Where("Uid = ?", Uid).Find(&users)
		if err1 != nil {
			return "null", err1
		}
	}
	ProfilePhot := users[0].ProfilePhot
	return ProfilePhot, nil
}
func GetSignature(Uid int64) (string, error) {

	var users []User
	if Uid != -1 {
		err1 := Db.Where("Uid = ?", Uid).Find(&users)
		if err1 != nil {
			return "null", err1
		}
	}
	Signature := users[0].Signature
	return Signature, nil
}

func GetLike(Uid int64) ([]Liked, error) {
	var likedBlogs []Liked
	var likeList []Like
	var blogLiked []Blog
	var reviews []Review
	var likeCount []Like
	var reviewAmount int64 = 0
	var likeAmount int64 = 0
	if Uid != -1 {
		err1 := Db.Where("Uid = ?", Uid).Find(&likeList)
		if err1 != nil {
			return []Liked{}, err1
		}
	}
	for _, liked := range likeList {
		Id := liked.Bid
		err2 := Db.Where("Id = ?", Id).Find(&blogLiked)
		if err2 != nil {
			return []Liked{}, err2
		}
		err3 := Db.Where("Bid = ?", Id).Find(&reviews)
		if err3 != nil {
			return []Liked{}, err3
		}
		for range reviews {
			reviewAmount += 1
		}
		err4 := Db.Where("Bid = ?", Id).Find(&likeCount)
		if err4 != nil {
			return []Liked{}, err4
		}
		for range likeCount {
			likeAmount += 1
		}
		likedBlog := Liked{Bid: blogLiked[0].Id, Content: blogLiked[0].Content, Picture: blogLiked[0].Picture, Title: blogLiked[0].Title, Visibility: blogLiked[0].Visibility, Time: blogLiked[0].Pub_time, Review: reviews, ReviewCount: reviewAmount, LikeCount: likeAmount}
		likedBlogs = append(likedBlogs, likedBlog)
	}
	return likedBlogs, nil
}
func GetCollect(c *gin.Context) ([]Collected, error) {
	var collectedBlogs []Collected
	var collectList []Collect
	var blogCollect []Blog
	var reviews []Review
	var collectCount []Collect
	var reviewAmount int64 = 0
	var collectAmount int64 = 0
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	if (Uid == -1) || (err != nil) {
		err1 := Db.Where("Uid = ?", Uid).Find(&collectList)
		if err1 != nil {
			return []Collected{}, err1
		}
	}
	for _, liked := range collectList {
		Id := liked.Bid
		err2 := Db.Where("Id = ?", Id).Find(&blogCollect)
		if err2 != nil {
			return []Collected{}, err2
		}
		err3 := Db.Where("Bid = ?", Id).Find(&reviews)
		if err3 != nil {
			return []Collected{}, err3
		}
		for range reviews {
			reviewAmount += 1
		}
		err4 := Db.Where("Bid = ?", Id).Find(&collectCount)
		if err4 != nil {
			return []Collected{}, err4
		}
		for range collectCount {
			collectAmount += 1
		}
		collectedBlog := Collected{Bid: blogCollect[0].Id, Content: blogCollect[0].Content, Picture: blogCollect[0].Picture, Title: blogCollect[0].Title, Visibility: blogCollect[0].Visibility, Time: blogCollect[0].Pub_time, Review: reviews, ReviewCount: reviewAmount, CollectCount: collectAmount}
		collectedBlogs = append(collectedBlogs, collectedBlog)
	}
	return collectedBlogs, nil
}
func GetJourney(c *gin.Context) (string, error) {

	var users []User
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	if (Uid == -1) || (err != nil) {
		err1 := Db.Where("Uid = ?", Uid).Find(&users)
		if err1 != nil {
			return "null", err1
		}
	}
	Signature := users[0].Signature
	return Signature, nil
}
func GetInfo(Uid int64) (User_detail, error) {
	var info User_detail
	err := Db.Where("uid = ?", Uid).Find(&info)
	if err != nil {
		return User_detail{}, err
	}
	return info, nil
}
