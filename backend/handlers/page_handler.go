package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

func ToIndex(c *gin.Context) {
	var Uid_tmp string
	if c.GetHeader("aToken") == "" {
		Uid_tmp = "-1"
	} else {
		aToken := c.GetHeader("aToken")
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	blogcount, _ := BlogCount(Uid)
	followcount, _ := FollowCount(Uid)
	followbycount, _ := FollowByCount(Uid)
	signature, _ := SignatureCheck(Uid)
	rank, _ := RankCheck(Uid)
	blogs, _ := BlogDisplay(Uid)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":     c.GetString("message"),
		"aToken":      c.GetHeader("aToken"),
		"rToken":      c.GetHeader("rToken"),
		"Uid":         Uid,
		"BlogCount":   blogcount,     //发帖数
		"FollowCount": followcount,   //关注数
		"FansCount":   followbycount, //粉丝数
		"Signature":   signature,     //个签
		"records":     rank,          //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		"blogs":       blogs,         //所有blog 用户名：.Author 头像：.Photo 发布时间：.Pub_time 可见性：.Visibility 内容：.Content 图片：.Picture 标题：.Title
	})
}

func ToIndex_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)

	fmt.Println(Uid)
	fmt.Println(aToken)

	isVip, _ := IsVip(Uid)
	blogcount, _ := BlogCount(Uid)
	followcount, _ := FollowCount(Uid)
	followbycount, _ := FollowByCount(Uid)
	signature, _ := SignatureCheck(Uid)
	rank, _ := RankCheck(Uid)
	blogs, _ := BlogDisplay(Uid)

	fmt.Println(blogs)
	c.JSON(http.StatusOK, gin.H{
		"message":     c.GetString("message"),
		"aToken":      c.GetHeader("aToken"),
		"rToken":      c.GetHeader("rToken"),
		"Uid":         Uid,
		"BlogCount":   blogcount,     //发帖数
		"FollowCount": followcount,   //关注数
		"FansCount":   followbycount, //粉丝数
		"Signature":   signature,     //个签
		"records":     rank,          //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		"blogs":       blogs,         //所有blog 用户名：.Author 头像：.Photo 发布时间：.Pub_time 可见性：.Visibility 内容：.Content 图片：.Picture 标题：.Title 是否被关注: .IsFollow
		"isVip":       isVip,         // 1是 or 0不是
	})
}
func ToRank_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	rank, _ := RankCheck(Uid)
	//tracks, _ := GetTracks()
	//todo: rank track
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"records": rank, //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		//"tracks":  tracks,
	})
}
func ToPublishTrack_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	tracks, _ := GetTracks(Uid)
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"tracks":  tracks, //.StrDate .StrTime .EndData .EndTime
	})
}
func ToPublish_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	type BlogApp struct {
		Title      string `json:"title"`
		Visibility string `json:"visibility"`
		Content    string `json:"content"`
		Tid        string `json:"route_id"`
		Picture    string `json:"imgs"`
	}
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	var blogApp BlogApp
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&blogApp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Visibility, _ := strconv.ParseInt(blogApp.Visibility, 10, 64)
	Tid, _ := strconv.ParseInt(blogApp.Tid, 10, 64)
	Bid := AddBlog(Uid, timeString, Visibility, blogApp.Content, blogApp.Picture, blogApp.Title, Tid)
	if Bid == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Blog upload unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog upload successfully."})
}

func ToLike_app(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	var user User
	var blog Blog
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	type LikeData struct {
		Bid string `json:"post_id"`
	}
	var likeData LikeData
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&likeData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	Bid, _ := strconv.ParseInt(likeData.Bid, 10, 64)
	Lid1 := AddLike(Uid, Bid, timeString)
	err1 := Db.Where("uid = ?", Uid).Find(&user)
	if err1 != nil {
		return
	}
	username := user.Name
	photo := user.ProfilePhot
	blogTitle := blog.Title
	err2 := Db.Where("bid = ?", Bid).Find(&blog)
	if err2 != nil {
		return
	}
	authorId := blog.Uid
	Lid2 := AddLikeApp(authorId, username, photo, blogTitle, timeString)
	if Lid1 == -1 || Lid2 == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Like unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like successfully."})
}

func ToCollect_app(c *gin.Context) {
	var user User
	var blog Blog
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	type CollectData struct {
		Bid string `json:"post_id"`
	}
	var collectData CollectData
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&collectData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Bid, _ := strconv.ParseInt(collectData.Bid, 10, 64)
	Cid1 := AddCollect(Uid, Bid, timeString)
	err1 := Db.Where("uid = ?", Uid).Find(&user)
	if err1 != nil {
		return
	}
	username := user.Name
	photo := user.ProfilePhot
	blogTitle := blog.Title
	err2 := Db.Where("bid = ?", Bid).Find(&blog)
	if err2 != nil {
		return
	}
	authorId := blog.Uid
	Cid2 := AddCollectApp(authorId, username, photo, blogTitle, timeString)
	if Cid1 == -1 || Cid2 == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Collect unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Collect successfully."})
}

func ToReview_app(c *gin.Context) {
	var user User
	var blog Blog
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	type ReviewData struct {
		Bid     string `json:"post_id"`
		Content string `json:"commentContent"`
	}
	var reviewData ReviewData
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&reviewData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	Bid, _ := strconv.ParseInt(reviewData.Bid, 10, 64)
	Rid1 := AddReview(Uid, Bid, timeString, reviewData.Content)
	err1 := Db.Where("uid = ?", Uid).Find(&user)
	if err1 != nil {
		return
	}
	username := user.Name
	photo := user.ProfilePhot
	blogTitle := blog.Title
	err2 := Db.Where("bid = ?", Bid).Find(&blog)
	if err2 != nil {
		return
	}
	authorId := blog.Uid
	Rid2 := AddReviewApp(authorId, username, photo, blogTitle, timeString)
	if Rid1 == -1 || Rid2 == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Review unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review successfully."})
}

func ToFollow_app(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	type FollowData struct {
		Uid string `json:"post_id"`
	}
	var followData FollowData
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&followData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	FollowUid, _ := strconv.ParseInt(followData.Uid, 10, 64)
	Fid := AddFollow(FollowUid, Uid, timeString)
	if Fid == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Follow unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Follow successfully."})
}

func ToLikeList_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	likelist, _ := GetLike(Uid)
	c.JSON(http.StatusOK, gin.H{
		"uid":      Uid,
		"aToken":   c.GetHeader("aToken"),
		"rToken":   c.GetHeader("rToken"),
		"likeList": likelist,
	})
}

func ToCollectList_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	collectlist, _ := GetCollect(Uid)
	c.JSON(http.StatusOK, gin.H{
		"message":     c.GetString("message"),
		"uid":         Uid,
		"aToken":      c.GetHeader("aToken"),
		"rToken":      c.GetHeader("rToken"),
		"collectList": collectlist,
	})
}

func ToVip_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	vip, _ := GetVip(Uid)
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"vip":     vip, // .StrTime .EndTime
	})
}

func TouploadInfo_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	type UserInfoData struct {
		Born     string `json:"born"`
		Status   string `json:"status"`
		Job      string `json:"job"`
		LivesIn  string `json:"livesin"`
		JoinedOn string `json:"joinedon"`
		Email    string `json:"email"`
	}
	var userInfoData UserInfoData
	w := c.Writer
	r := c.Request
	if err := json.NewDecoder(r.Body).Decode(&userInfoData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Did := AddUserInfo(Uid, userInfoData.Born, userInfoData.Status, userInfoData.Job, userInfoData.LivesIn, userInfoData.JoinedOn, userInfoData.Email)
	if Did == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "User Info update unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Info update successfully."})
}

func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-in.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-up.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToProfile(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	name, _ := GetName(Uid)
	email, _ := GetEmail(Uid)
	photo, _ := GetPhoto(Uid)
	motto, _ := GetSignature(Uid)
	likelist, _ := GetLike(Uid)
	journey, _ := GetTracks(Uid)
	blogs, _ := GetBlogs(Uid)
	c.HTML(http.StatusOK, "my-profile.html", gin.H{
		"message":  c.GetString("message"),
		"uid":      Uid,
		"aToken":   c.GetHeader("aToken"),
		"rToken":   c.GetHeader("rToken"),
		"username": name,
		"email":    email,
		"photo":    photo,
		"motto":    motto,
		"likeList": likelist,
		"journey":  journey,
		"blogs":    blogs,
	})
}
func ToProfile_app(c *gin.Context) {
	Uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)
	name, _ := GetName(Uid)
	email, _ := GetEmail(Uid)
	photo, _ := GetPhoto(Uid)
	motto, _ := GetSignature(Uid)
	likelist, _ := GetLike(Uid)
	journey, _ := GetTracks(Uid)
	Info, _ := GetInfo(Uid)
	blogs, _ := GetBlogs(Uid)
	//
	//fmt.Println("---------")
	//fmt.Println(Info)
	//fmt.Println("---------")
	//fmt.Println(c.Query("aToken"))
	//fmt.Println("---------")
	//fmt.Println(Info)
	//fmt.Println("---------")

	c.JSON(http.StatusOK, gin.H{
		"message":  c.GetString("message"),
		"uid":      Uid,
		"aToken":   c.GetHeader("aToken"),
		"rToken":   c.GetHeader("rToken"),
		"username": name,
		"email":    email,
		"photo":    photo,
		"motto":    motto,
		"likeList": likelist,
		"journey":  journey,
		"info":     Info,
		"blogs":    blogs,
	})
}
func ToVip(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToBill(c *gin.Context) {
	c.HTML(http.StatusOK, "checkout.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToDownload(c *gin.Context) {
	c.HTML(http.StatusOK, "app-download.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToHelp(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToHelpDetails(c *gin.Context, searchText string) {
	c.HTML(http.StatusOK, "help-details.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
	})
}

func ToMessages(c *gin.Context) {
	usernameCookie := http.Cookie{
		Name:  "username",
		Value: "zjy",
		Path:  "/",
	}
	http.SetCookie(c.Writer, &usernameCookie)
	c.HTML(http.StatusOK, "messaging.html", gin.H{
		"aToken": c.GetHeader("aToken"),
		"rToken": c.GetHeader("rToken"),
	})
	//Claim, _ := CheckToken(c.GetString)
	//Uid := Claim.Uid

}

func ToNotifications(c *gin.Context) {
	c.HTML(http.StatusOK, "notifications.html", gin.H{
		"aToken": c.GetHeader,
		"rToken": c.GetHeader("rToken"),
	})
}

func ToSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"aToken": c.GetHeader("aToken"),
		"rToken": c.GetHeader("rToken"),
	})
}

func ToWs(c *gin.Context) {
	wshandler(c.Writer, c.Request)
}

func ToPostDetails(Id string, c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	Bid, _ := strconv.ParseInt(Id, 10, 64)
	blog, _ := GetBlog(Bid)
	track, _ := GetTrack(blog.Tid)
	reviews, _ := GetReviews(Bid)
	photo, _ := GetPhoto(Uid)
	c.HTML(http.StatusOK, "post-details.html", gin.H{
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"photo":   photo,
		"blog":    blog,
		"track":   track,
		"reviews": reviews, //评论者.Reviewer 评论者头像.Reviewer_photo 时间.Time 内容.Content
	})
}
func ToLog(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	photo, _ := GetPhoto(Uid)
	c.JSON(http.StatusOK, gin.H{
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"photo":   photo,
		"browser": GetBrowser(c),
	})
}

func ToPostDetails_app(Id string, c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	Bid, _ := strconv.ParseInt(Id, 10, 64)
	blog, _ := GetBlog(Bid)
	track, _ := GetTrack(blog.Tid)
	reviews, _ := GetReviews(Bid)
	photo, _ := GetPhoto(Uid)
	c.JSON(http.StatusOK, gin.H{
		"aToken":  c.GetHeader("aToken"),
		"rToken":  c.GetHeader("rToken"),
		"photo":   photo,
		"blog":    blog,
		"track":   track,
		"reviews": reviews, //评论者.Reviewer 评论者头像.Reviewer_photo 时间.Time 内容.Content
	})
}
func ToSearchedBlogs(c *gin.Context) {
	text := c.Param("text")
	blogDisplays, err := SearchBlogs(text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogDisplays)
}

func ToFriendsBlogs(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)

	blogDisplays, err := GetFriendsBlogs(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogDisplays)
}

func ToShowLiked_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)

	likedMess, err := GetLikedMess(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likedMess)
}

func ToShowReviewed_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)

	reviewedMess, err := GetReviewedMess(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviewedMess)
}
func ToShowCollected_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	var Uid_tmp string
	if aToken == "" {
		Uid_tmp = "-1"
	} else {
		Claim, err := CheckToken(aToken)
		if err != nil {
			Uid_tmp = "-1"
		} else {
			Uid_tmp = Claim.Uid
		}
	}
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)

	collectedMess, err := GetCollectedMess(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, collectedMess)
}
func ToFollowing(c *gin.Context) {
	Uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)

	followings, err := GetFollowings(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, followings)
}
func ToFans(c *gin.Context) {
	Uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)

	fans, err := GetFans(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fans)
}

func ToFriends(c *gin.Context) {
	Uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)

	friends, err := GetFriends(Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func ToAllUsers(c *gin.Context) {
	users := AllUsers()
	c.JSON(http.StatusOK, users)
}
