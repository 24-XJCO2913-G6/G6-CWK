package handlers

import (
	. "main/backend/models"
	"sort"
	"strconv"
)

func BlogCount(Uid int64) (int64, error) {
	var blogList []Blog
	var TotalamountBlog int64
	TotalamountBlog = 0
	if Uid != -1 {
		err := Db.Where("uid = ?", Uid).Find(&blogList)
		if err != nil {
			return 0, err
		}
		TotalamountBlog = int64(len(blogList))
	}

	return TotalamountBlog, nil
}

func FollowByCount(Uid int64) (int64, error) {
	var followByList []Follow
	var TotalamountFollowBy int64
	TotalamountFollowBy = 0

	if Uid != -1 {
		err := Db.Where("follow_id = ?", Uid).Find(&followByList)
		if err != nil {
			return 0, err
		}
		TotalamountFollowBy = int64(len(followByList))
	}

	return TotalamountFollowBy, nil
}

func FollowCount(Uid int64) (int64, error) {
	var followList []Follow
	var TotalamountFollow int64 = 0
	if Uid != -1 {
		err := Db.Where("follow_by_id = ?", Uid).Find(&followList)
		if err != nil {
			return 0, err
		}
		TotalamountFollow = int64(len(followList))
	}

	return TotalamountFollow, nil
}

func SignatureCheck(Uid int64) (string, error) {
	var signaturelist []User
	var Signature string = ""
	if Uid != -1 {
		err := Db.Where("uid = ?", Uid).Find(&signaturelist)
		if err != nil {
			return "", err
		}
		if len(signaturelist) != 0 {
			Signature = signaturelist[0].Signature
		}
	}

	return Signature, nil
}

func RankCheck(Uid int64) ([]Records, error) {
	var recordlist []Records
	var tracks []Track
	var totaldis float64 = 0

	// 查询并按照 i 值进行分组，并将 distance 求和
	users, _ := GetFriends(Uid)
	_ = append(users, Friend{Uid: Uid})
	for _, user := range users {
		// 处理每一行数据
		tracks, _ = GetTracks(user.Uid)
		if len(tracks) != 0 {
			maxTrack := tracks[0].Coordinates
			maxDis := tracks[0].Distance
			for _, record := range tracks {
				Dis, _ := strconv.ParseFloat(record.Distance, 64)
				totaldis += Dis
				if maxDis < record.Distance {
					maxDis = record.Distance
					maxTrack = record.Coordinates
				}
			}
			DisSum := strconv.FormatFloat(totaldis, 'f', -1, 64)
			record := Records{Uid: user.Uid, Distance: DisSum, Name: user.Name, Photo: user.Pic, Coordinates: maxTrack}
			recordlist = append(recordlist, record)
		}

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
	var track Track
	var followed Follow
	var isFollow int64
	err2 := Db.Find(&blogs)
	if err2 != nil {
		return []Blog_display{}, err2
	}
	if len(blogs) != 0 {
		for _, blog := range blogs {
			Uid := blog.Uid
			err1 := Db.Where("uid = ?", Uid).Find(&authors)
			if err1 != nil {
				return []Blog_display{}, err1
			}
			has, err := Db.Where("follow_by_id = ? AND follow_id = ?", Uid, authors[0].Uid).Get(&followed)
			if err != nil {
				return []Blog_display{}, err
			} else if !has {
				isFollow = 0
			} else {
				isFollow = 1
			}
			err4 := Db.Where("tid = ?", blog.Tid).Find(&track)
			if err4 != nil {
				return []Blog_display{}, err4
			}
			author := authors[0].Name
			photo := authors[0].ProfilePhot
			Aid := authors[0].Uid
			blog_dis := Blog_display{Author: author, Photo: photo, Pub_time: blog.Pub_time, Visibility: blog.Visibility,
				Content: blog.Content, Picture: blog.Picture, Title: blog.Title, IsFollow: isFollow, Coordinates: track.Coordinates, AuthorId: Aid}
			blogs_dis = append(blogs_dis, blog_dis)
		}
	}

	return []Blog_display{}, nil
}
