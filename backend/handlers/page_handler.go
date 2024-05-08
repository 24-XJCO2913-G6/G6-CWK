package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

func ToIndex(c *gin.Context) {
	//err := c.Request.ParseForm()
	//if err != nil {
	//	return
	//}
	var Uid string
	if c.PostForm("aToken") == "" {
		Uid = "-1"
	} else {
		aToken := c.PostForm("aToken")
		Claim, _ := CheckToken(aToken)
		Uid = Claim.Uid
	}
	//blogcount, _ := BlogCount(c)
	//followcount, _ := FollowCount(c)
	//followbycount, _ := FollowByCount(c)
	//signature, _ := SignatureCheck(c)
	//rank, _ := RankCheck()
	//blogs, _ := BlogDisplay()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": c.PostForm("message"),
		"aToken":  c.PostForm("aToken"),
		"rToken":  c.PostForm("rToken"),
		"Uid":     Uid,
		//"BlogCount":   blogcount,     //发帖数
		//"FollowCount": followcount,   //关注数
		//"FansCount":   followbycount, //粉丝数
		//"Signature":   signature,     //个签
		//"records":     rank,          //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		//"blogs":       blogs,         //所有blog 用户名：.Author 头像：.Photo 发布时间：.Pub_time 可见性：.Visibility 内容：.Content 图片：.Picture 标题：.Title
	})
}

func ToIndex_app(c *gin.Context) {
	var Uid string
	Claim, res := CheckToken(c.GetString("aToken"))
	if res != nil {
		Uid = "-1"
	}
	Uid = Claim.Uid
	blogcount, _ := BlogCount(c)
	followcount, _ := FollowCount(c)
	followbycount, _ := FollowByCount(c)
	signature, _ := SignatureCheck(c)
	rank, _ := RankCheck()
	blogs, _ := BlogDisplay()
	c.JSON(http.StatusOK, gin.H{
		"message":     c.GetString("message"),
		"aToken":      c.GetString("aToken"),
		"rToken":      c.GetString("rToken"),
		"Uid":         Uid,
		"BlogCount":   blogcount,     //发帖数
		"FollowCount": followcount,   //关注数
		"FansCount":   followbycount, //粉丝数
		"Signature":   signature,     //个签
		"records":     rank,          //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		"blogs":       blogs,         //所有blog 用户名：.Author 头像：.Photo 发布时间：.Pub_time 可见性：.Visibility 内容：.Content 图片：.Picture 标题：.Title
	})
}
func ToRank_app(c *gin.Context) {
	rank, _ := RankCheck()
	//tracks, _ := GetTracks()
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"records": rank, //排名(降序) 距离：.Distance 用户名：.Name 头像：.Photo
		//"tracks":  tracks,
	})
}
func ToPublishTrack_app(c *gin.Context) {
	aToken := c.GetHeader("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
	tracks, _ := GetTracks(Uid)
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"tracks":  tracks, //.StrDate .StrTime .EndData .EndTime
	})
}
func ToPublish_app(c *gin.Context) {
	//aToken := c.GetString("aToken")
	//Claims, _ := CheckToken(aToken)
	//Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
	//currentTime := time.Now()
	//timeString := currentTime.Format("2006-01-02 15:04:05")
	//Title := c.PostForm("title")
	//Visibility_tmp := c.PostForm("visibility")
	//Visibility, _ := strconv.ParseInt(Visibility_tmp, 10, 64)
	//Content := c.PostForm("content")
	//Tid_tmp := c.PostForm("route_id")
	//Tid, _ := strconv.ParseInt(Tid_tmp, 10, 64)
	//Picture := c.PostForm("imgs")
	//Bid := AddBlog(Uid, timeString, Visibility, Content, Picture, Title, Tid)
	//if Bid == -1 {
	//	c.JSON(http.StatusOK, gin.H{"error": "Blog upload unsuccessfully."})
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Blog upload successfully."})

	aToken := c.GetHeader("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
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
	aToken := c.GetHeader("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
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
	Lid := AddLike(Uid, Bid, timeString)

	if Lid == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Like unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like successfully."})
}

func ToCollect_app(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	aToken := c.GetHeader("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	Bid_tmp := c.PostForm("post_id")
	Bid, _ := strconv.ParseInt(Bid_tmp, 10, 64)
	Cid := AddCollect(Uid, Bid, timeString)
	if Cid == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Collect unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Collect successfully."})
}

func ToFollow_app(c *gin.Context) {
	// 根据请求携带的aToken判断添加路径的用户
	aToken := c.GetHeader("aToken")
	Claims, _ := CheckToken(aToken)
	Uid, _ := strconv.ParseInt(Claims.Uid, 10, 64)
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	FollowUid_tmp := c.PostForm("user_id")
	FollowUid, _ := strconv.ParseInt(FollowUid_tmp, 10, 64)
	Fid := AddFollow(FollowUid, Uid, timeString)
	if Fid == -1 {
		c.JSON(http.StatusOK, gin.H{"error": "Follow unsuccessfully."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Follow successfully."})
}

func ToLikeList_app(c *gin.Context) {
	Claim, _ := CheckToken(c.GetHeader("aToken"))
	Uid_tmp := Claim.Uid
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	likelist, _ := GetLike(c)
	c.JSON(http.StatusOK, gin.H{
		"uid":      Uid,
		"message":  c.GetString("message"),
		"aToken":   c.GetString("aToken"),
		"rToken":   c.GetString("rToken"),
		"likeList": likelist,
	})
}

func ToCollectList_app(c *gin.Context) {
	Claim, _ := CheckToken(c.GetHeader("aToken"))
	Uid_tmp := Claim.Uid
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	collectlist, _ := GetCollect(c)
	c.JSON(http.StatusOK, gin.H{
		"uid":         Uid,
		"message":     c.GetString("message"),
		"aToken":      c.GetString("aToken"),
		"rToken":      c.GetString("rToken"),
		"collectList": collectlist,
	})
}

func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-in.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-up.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToProfile(c *gin.Context) {
	Claim, _ := CheckToken(c.GetHeader("aToken"))
	Uid_tmp := Claim.Uid
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	name, _ := GetName(c)
	email, _ := GetEmail(c)
	photo, _ := GetPhoto(c)
	motto, _ := GetSignature(c)
	likelist, _ := GetLike(c)
	journey, _ := GetTracks(Uid)
	c.HTML(http.StatusOK, "my-profile.html", gin.H{
		"uid":      Uid,
		"message":  c.GetString("message"),
		"aToken":   c.GetString("aToken"),
		"rToken":   c.GetString("rToken"),
		"username": name,
		"email":    email,
		"photo":    photo,
		"motto":    motto,
		"likeList": likelist,
		"journey":  journey,
	})
}
func ToProfile_app(c *gin.Context) {
	Claim, _ := CheckToken(c.GetString("aToken"))
	Uid_tmp := Claim.Uid
	Uid, _ := strconv.ParseInt(Uid_tmp, 10, 64)
	name, _ := GetName(c)
	email, _ := GetEmail(c)
	photo, _ := GetPhoto(c)
	motto, _ := GetSignature(c)
	likelist, _ := GetLike(c)
	journey, _ := GetTracks(Uid)
	c.JSON(http.StatusOK, gin.H{
		"uid":      Uid,
		"message":  c.GetString("message"),
		"aToken":   c.GetString("aToken"),
		"rToken":   c.GetString("rToken"),
		"username": name,
		"email":    email,
		"photo":    photo,
		"motto":    motto,
		"likeList": likelist,
		"journey":  journey,
	})
}
func ToVip(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToBill(c *gin.Context) {
	c.HTML(http.StatusOK, "checkout.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToDownload(c *gin.Context) {
	c.HTML(http.StatusOK, "app-download.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToHelp(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToHelpDetails(c *gin.Context, searchText string) {
	c.HTML(http.StatusOK, "help-details.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
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
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
	//Claim, _ := CheckToken(c.GetString("aToken"))
	//Uid := Claim.Uid

}

func ToNotifications(c *gin.Context) {
	c.HTML(http.StatusOK, "notifications.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
	})
}

func ToWs(c *gin.Context) {
	wshandler(c.Writer, c.Request)
}

func ToPostDetails(Id string, c *gin.Context) {
	Bid, _ := strconv.ParseInt(Id, 10, 64)
	blog, _ := GetBlog(Bid)
	track, _ := GetTrack(blog.Tid)
	reviews, _ := GetReviews(Bid)
	photo, _ := GetPhoto(c)
	c.HTML(http.StatusOK, "post-details.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"photo":   photo,
		"blog":    blog,
		"track":   track,
		"reviews": reviews, //评论者.Reviewer 评论者头像.Reviewer_photo 时间.Time 内容.Content
	})
}
func ToLog(c *gin.Context) {
	photo, _ := GetPhoto(c)
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"photo":   photo,
		"browser": GetBrowser(c),
	})
}

func ToPostDetails_app(Id string, c *gin.Context) {
	Bid, _ := strconv.ParseInt(Id, 10, 64)
	blog, _ := GetBlog(Bid)
	track, _ := GetTrack(blog.Tid)
	reviews, _ := GetReviews(Bid)
	photo, _ := GetPhoto(c)
	c.JSON(http.StatusOK, gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"photo":   photo,
		"blog":    blog,
		"track":   track,
		"reviews": reviews, //评论者.Reviewer 评论者头像.Reviewer_photo 时间.Time 内容.Content
	})
}
