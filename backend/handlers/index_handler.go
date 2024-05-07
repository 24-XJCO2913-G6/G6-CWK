package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	. "main/backend/models"
	"sort"
	"strconv"
)

func BlogCount(c *gin.Context) (int64, error) {

	var blogList []Blog
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
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
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
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
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
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
	aToken := c.GetString("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, err := strconv.ParseInt(Claims.Uid, 10, 64)
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

func RankCheck() ([]Records, error) {
	var recordlist []Records
	var tracks []Track
	var totaldis float64 = 0
	// 查询并按照 i 值进行分组，并将 distance 求和
	var users []User
	err := Db.Find(&users)
	if err != nil {
		log.Fatalf("Failed to query database: %v", err)
	}

	// 遍历结果集
	for _, user := range users {
		// 处理每一行数据
		tracks, err = GetTracks(user.Uid)
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
		blog_dis := Blog_display{Author: author, Photo: photo, Pub_time: blog.Pub_time, Visibility: blog.Visibility, Content: blog.Content, Picture: blog.Picture, Tag: blog.Tag}
		blogs_dis = append(blogs_dis, blog_dis)
	}
	return []Blog_display{}, nil
}
