package handlers

import (
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
	return user.Signature, nil
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
		return nil, nil
	}

	err := Db.Find(&likeList)
	if err != nil {
		return nil, nil
	}
	if len(likeList) != 0 {
		for _, liked := range likeList {
			Id := liked.Bid
			err2 := Db.Where("id = ?", Id).Find(&blogLiked)
			if err2 != nil {
				return nil, err2
			}
			err3 := Db.Where("bid = ?", Id).Find(&reviews)
			if err3 != nil {
				return nil, err3
			}
			err4 := Db.Where("bid = ?", Id).Find(&likeCount)
			if err4 != nil {
				return nil, err4
			}
			reviewAmount = int64(len(reviews))
			likeAmount = int64(len(likeCount))
			if len(blogLiked) != 0 {
				likedBlog := Liked{Bid: (blogLiked)[0].Id, Content: (blogLiked)[0].Content, Picture: (blogLiked)[0].Picture, Title: (blogLiked)[0].Title, Visibility: (blogLiked)[0].Visibility, Time: (blogLiked)[0].Pub_time, Review: reviews, ReviewCount: reviewAmount, LikeCount: likeAmount}
				likedBlogs = append(likedBlogs, likedBlog)
			}
		}
	}

	return likedBlogs, nil

}

func GetCollect(Uid int64) ([]Blog_display, error) {
	var collectList []Collect
	var blogs []Blog_display
	if Uid == -1 {
		return nil, nil
	}
	blogCollect, _ := BlogDisplay(Uid)
	_ = Db.Where("uid = ?", Uid).Find(&collectList)
	for _, liked := range collectList {
		for _, blog := range blogCollect {
			if blog.Bid == liked.Bid {
				blogs = append(blogs, blog)
			}
		}
	}

	//fmt.Println("------------collection--------------")
	//fmt.Println(blogs)
	//fmt.Println("------------collection--------------")

	return blogs, nil
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
	var infos []User_detail
	err := Db.Where("uid = ?", Uid).Find(&infos)
	if err != nil {
		return User_detail{}, nil
	}
	if len(infos) == 0 {
		return User_detail{}, nil
	}
	//fmt.Println(infos)
	return infos[0], nil
}

func GetBlogs(Uid int64) ([]Blog, error) {
	var blogs []Blog
	err := Db.Where("uid = ?", Uid).Find(&blogs)
	if err != nil {
		return []Blog{}, nil
	}
	if len(blogs) == 0 {
		return []Blog{}, nil
	}
	return blogs, nil
}
