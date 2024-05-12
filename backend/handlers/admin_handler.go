package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	CreatedTime   string `json:"created_time"`
	IsVIP         bool   `json:"is_vip"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	VipExpiryDate string `json:"vip_expiry_date"`
	Photo         string `json:"photo"`
	Birthday      string `json:"birthday"`
	Status        string `json:"status"`
	Job           string `json:"job"`
	Hometown      string `json:"hometown"`
}

type PendingBlog struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	PostTime  string `json:"post_time"`
	PostTitle string `json:"post_title"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TrackID   int64  `json:"track_id"`
	PathData  string `json:"path_data"`
}

func commonPrefix(str1, str2 string) string {
	// 找到两个字符串中较短的那个字符串的长度
	minLength := len(str1)
	if len(str2) < minLength {
		minLength = len(str2)
	}

	// 遍历两个字符串，找出相同的前缀
	i := 0
	for i < minLength && str1[i] == str2[i] {
		i++
	}

	return str1[:i]
}

func ToIndexAdmin(c *gin.Context) {
	var Income float64 = 0.0
	var orders []Order
	err := Db.Find(&orders)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
	}
	if len(orders) > 0 {
		for _, order := range orders {
			money, _ := strconv.ParseFloat(order.Price, 64)
			Income += money
		}
	}

	TodayUsers := 0
	TodayBlogs := 0

	users := AllUsers()
	var blogs []Blog
	err = Db.Find(&blogs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
	}
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	for _, user := range users {
		if len(commonPrefix(date, user.CreatedTime)) == len(date) {
			TodayUsers += 1
		}
	}
	for _, blog := range blogs {
		if len(commonPrefix(date, blog.Pub_time)) == len(date) {
			TodayBlogs += 1
		}
	}

	TotalUsers := len(users)
	TotalBlogs := len(blogs)

	var tracks []Track
	err = Db.Find(&tracks)
	if err != nil || len(tracks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	drivingCount := 0
	runningCount := 0
	cyclingCount := 0

	for _, track := range tracks {
		if track.Mode == "driving" {
			drivingCount += 1
		} else if track.Mode == "cycling" {
			cyclingCount += 1
		} else {
			runningCount += 1

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"totalIncome":   Income,
		"todayNewUsers": TodayUsers,
		"totalUsers":    TotalUsers,
		"todayPosts":    TodayBlogs,
		"totalPosts":    TotalBlogs,
		"drivingCount":  drivingCount,
		"runningCount":  runningCount,
		"cyclingCount":  cyclingCount,
	})
}

func ToUsersInfo(c *gin.Context) {
	users := AllUsers()

	var userInfos []UserInfo

	for _, user := range users {
		var vip_expiry_date string
		if user.IsVip {
			var vipInfo []Vip
			err := Db.Where("uid = ?", user.Uid).Find(&vipInfo)
			if err != nil || len(vipInfo) == 0 {
				vip_expiry_date = ""
			} else {
				vip_expiry_date = vipInfo[0].EndTime
			}
		} else {
			vip_expiry_date = ""
		}

		var userDetail []User_detail
		err := Db.Where("uid = ?", user.Uid).Find(&userDetail)
		var tmpUserDetail = User_detail{
			Did:      0,
			Uid:      0,
			Born:     "",
			Status:   "",
			Job:      "",
			LivesIn:  "",
			JoinedOn: "",
			Email:    "",
		}
		if err == nil && len(userDetail) != 0 {
			tmpUserDetail = userDetail[0]
		}
		userInfo := UserInfo{
			ID:            user.Uid,
			Name:          user.Name,
			CreatedTime:   user.CreatedTime,
			IsVIP:         user.IsVip,
			Password:      user.Passwd,
			Email:         user.Email,
			VipExpiryDate: vip_expiry_date,
			Photo:         user.ProfilePhot,
			Birthday:      tmpUserDetail.Born,
			Status:        tmpUserDetail.Status,
			Job:           tmpUserDetail.Job,
			Hometown:      tmpUserDetail.LivesIn,
		}

		userInfos = append(userInfos, userInfo)
	}
	c.JSON(http.StatusOK, userInfos)
}

func ToOrdersInfo(c *gin.Context) {
	var orders []Order
	err := Db.Find(&orders)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
	}
	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"orders": []Order{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func ToPendingsInfo(c *gin.Context) {
	var blogs []VipBlog
	var pendingBlogs []PendingBlog

	err := Db.Find(&blogs)
	if err != nil || len(blogs) == 0 {
		c.JSON(http.StatusOK, gin.H{"pendings": []PendingBlog{}})
	}

	for _, blog := range blogs {
		var pending_blog = PendingBlog{
			ID:        blog.Vid,
			Name:      "",
			PostTime:  blog.Pub_time,
			PostTitle: blog.Title,
			Title:     blog.Title,
			Content:   blog.Content,
			TrackID:   blog.Tid,
			PathData:  "",
		}

		user := &User{Uid: blog.Uid}
		track := &Track{Tid: blog.Tid}
		has, err := Db.Get(user)
		if has && err == nil {
			pending_blog.Name = user.Name
		}

		has, err = Db.Get(track)
		if has && err == nil {
			pending_blog.PathData = track.Coordinates
		}

		pendingBlogs = append(pendingBlogs, pending_blog)
	}

	c.JSON(http.StatusOK, pendingBlogs)
}

func ToLogsInfo(c *gin.Context) {
	logs, err := GetLogs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"logs": logs})
	}
}

func ToPassVipBlog(c *gin.Context) {
	Vid_tmp := c.Param("Vid")
	Vid, _ := strconv.ParseInt(Vid_tmp, 10, 64)

	meg := PassVipBlog(Vid)
	state := 200
	if meg == "Pass failed" {
		state = 400
	}
	c.JSON(state, gin.H{"message": meg})
}

func ToDeleteVipBlog(c *gin.Context) {
	Vid_tmp := c.Param("Vid")
	Vid, _ := strconv.ParseInt(Vid_tmp, 10, 64)

	meg := DeleteVipBlog(Vid)
	state := 200
	if meg == "Delete failed" {
		state = 400
	}
	c.JSON(state, gin.H{"message": meg})
}

func Dashboard(c *gin.Context) {
	IncomeW, err := IncomeEstWeek()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	IncomeM, err := IncomeEstMonth()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	IncomeY, err := IncomeEstYear()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"week":  IncomeW,
		"month": IncomeM,
		"year":  IncomeY})
}
