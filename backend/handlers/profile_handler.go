package handlers

import (
	"fmt"
	. "main/backend/models"
)

func GetName(Uid int64) (string, error) {

	var user = &User{Uid: Uid}
	if Uid == -1 {
		return "null", nil
	}
	has, _ := Db.Get(user)
	if !has {
		return "null", nil
	}
	fmt.Println(user)
	return user.Name, nil
}

func GetEmail(Uid int64) (string, error) {

	var user = &User{Uid: Uid}
	if Uid == -1 {
		return "null", nil
	}
	has, _ := Db.Get(user)
	if !has {
		return "null", nil
	}
	return user.Email, nil
}

func GetPhoto(Uid int64) (string, error) {

	var user = &User{Uid: Uid}
	if Uid == -1 {
		return "null", nil
	}
	has, _ := Db.Get(user)
	if !has {
		return "null", nil
	}
	return user.ProfilePhot, nil
}
func GetSignature(Uid int64) (string, error) {

	var user = &User{Uid: Uid}
	if Uid == -1 {
		return "null", nil
	}
	has, _ := Db.Get(user)
	if !has {
		return "null", nil
	}
	return user.Name, nil
}

func GetLike(Uid int64) ([]Liked, error) {
	var likedBlogs = new([]Liked)
	var likeList = new([]Like)
	var blogLiked = new([]Blog)
	var reviews = new([]Review)
	var likeCount = new([]Like)
	var reviewAmount int64 = 0
	var likeAmount int64 = 0
	if Uid != -1 {
		return nil, nil
	}
	has, _ := Db.Get(likeList)
	if !has {
		return nil, nil
	}
	for _, liked := range *likeList {
		Id := liked.Bid
		err2 := Db.Where("Id = ?", Id).Find(blogLiked)
		if err2 != nil {
			return nil, err2
		}
		err3 := Db.Where("Bid = ?", Id).Find(reviews)
		if err3 != nil {
			return nil, err3
		}
		err4 := Db.Where("Bid = ?", Id).Find(likeCount)
		if err4 != nil {
			return nil, err4
		}
		reviewAmount = int64(len(*reviews))
		likeAmount = int64(len(*likeCount))
		likedBlog := Liked{Bid: (*blogLiked)[0].Id, Content: (*blogLiked)[0].Content, Picture: (*blogLiked)[0].Picture, Title: (*blogLiked)[0].Title, Visibility: (*blogLiked)[0].Visibility, Time: (*blogLiked)[0].Pub_time, Review: *reviews, ReviewCount: reviewAmount, LikeCount: likeAmount}
		*likedBlogs = append(*likedBlogs, likedBlog)
	}
	return *likedBlogs, nil

}
func GetCollect(Uid int64) ([]Collected, error) {
	var collectedBlogs = new([]Collected)
	var collectList = new([]Collect)
	var blogCollect = new([]Blog)
	var reviews = new([]Review)
	var collectCount = new([]Collect)
	var reviewAmount int64 = 0
	var collectAmount int64 = 0
	if Uid == -1 {
		return nil, nil
	}
	for _, liked := range *collectList {
		Id := liked.Bid
		err2 := Db.Where("Id = ?", Id).Find(blogCollect)
		if err2 != nil {
			return nil, err2
		}
		err3 := Db.Where("Bid = ?", Id).Find(reviews)
		if err3 != nil {
			return nil, err3
		}
		reviewAmount = int64(len(*reviews))
		err4 := Db.Where("Bid = ?", Id).Find(collectCount)
		if err4 != nil {
			return []Collected{}, err4
		}
		collectAmount = int64(len(*collectCount))
		collectedBlog := Collected{Bid: (*blogCollect)[0].Id, Content: (*blogCollect)[0].Content, Picture: (*blogCollect)[0].Picture, Title: (*blogCollect)[0].Title, Visibility: (*blogCollect)[0].Visibility, Time: (*blogCollect)[0].Pub_time, Review: (*reviews), ReviewCount: reviewAmount, CollectCount: collectAmount}
		*collectedBlogs = append(*collectedBlogs, collectedBlog)
	}
	return *collectedBlogs, nil
}

//func GetJourney(c *gin.Context) (string, error) {
//
//	var users []User
//	aToken := c.GetHeader("aToken")
//	var Uid_tmp string
//	if aToken == "" {
//		Uid_tmp = "-1"
//	} else {
//		Claim, _ := CheckToken(aToken)
//		Uid_tmp = Claim.Uid
//	}
//	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
//	if (Uid == -1) || (err != nil) {
//		err1 := Db.Where("Uid = ?", Uid).Find(&users)
//		if err1 != nil {
//			return "null", err1
//		}
//	}
//	Signature := users[0].Signature
//	return Signature, nil
//}

func GetInfo(Uid int64) (User_detail, error) {
	var info = &User_detail{Uid: Uid}
	has, err := Db.Get(&info)
	if has {
		return User_detail{}, err
	}
	return *info, nil
}
