package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"sort"
	"strconv"
)

func BlogCount(c *gin.Context) (int64, error) {

	var blogList []Blog
	var Uid_tmp string
	if c.PostForm("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	var TotalamountBlog int64
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&blogList)
		if err != nil {
			return 0, err
		}

		for range blogList {
			TotalamountBlog += 1
		}
	}

	return TotalamountBlog, nil
}

func FollowByCount(c *gin.Context) (int64, error) {
	var followByList []Follow
	var Uid_tmp string
	if c.PostForm("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	var TotalamountFollowBy int64
	if (Uid == -1) || (err != nil) {
		err := Db.Where("FollowId = ?", Uid).Find(&followByList)
		if err != nil {
			return 0, err
		}
		for range followByList {
			TotalamountFollowBy += 1
		}
	}

	return TotalamountFollowBy, nil
}

func FollowCount(c *gin.Context) (int64, error) {
	var followList []Follow
	var Uid_tmp string
	if c.PostForm("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	var TotalamountFollow int64 = 0
	if (Uid == -1) || (err != nil) {
		err := Db.Where("FollowById = ?", Uid).Find(&followList)
		if err != nil {
			return 0, err
		}

		for range followList {
			TotalamountFollow += 1
		}
	}

	return TotalamountFollow, nil
}

func SignatureCheck(c *gin.Context) (string, error) {
	var signaturelist []User
	var Uid_tmp string
	if c.PostForm("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, err := strconv.ParseInt(Uid_tmp, 10, 64)
	var Signature string
	if (Uid == -1) || (err != nil) {
		err := Db.Where("Uid = ?", Uid).Find(&signaturelist)
		if err != nil {
			return "null", err
		}
		Signature = signaturelist[0].Signature
	}

	return Signature, nil
}

func RankCheck(c *gin.Context) ([]Records, error) {
	var recordlist []Records
	var tracks []Track
	var followlist []Follow
	var followbylist []Follow
	arr := make([]int64, 0)
	var totaldis float64 = 0
	// 查询并按照 i 值进行分组，并将 distance 求和
	var users []User
	var friend User
	var Uid_tmp string
	if c.PostForm("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid_tmp = Claim.Uid
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	err := Db.Where("FollowById = ?", Uid).Find(&followlist)
	for _, follower := range followlist {
		has, err := Db.Where("FollowById = ? AND FollowId = ?", follower.FollowId, Uid).Get(&followbylist)

		if err != nil {
			return []Records{}, err
		} else if !has {
			continue
		} else {
			arr = append(arr, follower.FollowId)
		}

	}
	if err != nil {
		return []Records{}, err
	}
	for _, friendid := range arr {
		has, err := Db.Where("Uid = ?", friendid).Get(&friend)
		if err != nil {
			return []Records{}, err
		} else if !has {
			fmt.Printf("user(id) not exist:%d\n", friendid)
		} else {
			users = append(users, friend)
		}
	}
	// 遍历结果集
	for _, user := range users {
		// 处理每一行数据
		tracks, _ = GetTracks(user.Uid)
		for _, record := range tracks {
			Dis, _ := strconv.ParseFloat(record.Distance, 64)
			totaldis += Dis
		}
		DisSum := strconv.FormatFloat(totaldis, 'f', -1, 64)
		record := Records{Uid: user.Uid, Distance: DisSum, Name: user.Name, Photo: user.ProfilePhot}
		recordlist = append(recordlist, record)
	}
	sort.Slice(recordlist, func(i, j int) bool {
		return recordlist[i].Distance > recordlist[j].Distance
	})
	// 输出结果

	return recordlist, nil
}

func BlogDisplay() ([]Blog_display, error) {
	var blogs []Blog
	var authors []User
	var blogs_dis []Blog_display
	err := Db.Find(&blogs)
	if err != nil {
		return []Blog_display{}, err
	}
	for _, blog := range blogs {
		Uid := blog.Uid
		err := Db.Where("Uid = ?", Uid).Find(&authors)
		if err != nil {
			return []Blog_display{}, err
		}
		author := authors[0].Name
		photo := authors[0].ProfilePhot
		blog_dis := Blog_display{Author: author, Photo: photo, Pub_time: blog.Pub_time, Visibility: blog.Visibility, Content: blog.Content, Picture: blog.Picture, Title: blog.Title}
		blogs_dis = append(blogs_dis, blog_dis)
	}
	return []Blog_display{}, nil
}
