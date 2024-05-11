package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
	"strconv"
	"time"
)

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

	c.JSON(http.StatusOK, gin.H{
		"totalIncome":   Income,
		"todayNewUsers": TodayUsers,
		"totalUsers":    TotalUsers,
		"todayPosts":    TodayBlogs,
		"totalPosts":    TotalBlogs,
	})
}

func ToUsersInfo(c *gin.Context) {
	users := AllUsers()

	type UserInfo struct {
		id              int64
		name            string
		created_time    string
		is_vip          bool
		password        string
		email           string
		vip_expiry_date string
		photo           string
		birthday        string
		status          string
		job             string
		hometown        string
	}
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
		if err != nil || len(userDetail) == 0 {
			tmpUserDetail = userDetail[0]
		}
		userInfo := UserInfo{
			id:              user.Uid,
			name:            user.Name,
			created_time:    user.CreatedTime,
			is_vip:          user.IsVip,
			password:        user.Passwd,
			email:           user.Email,
			vip_expiry_date: vip_expiry_date,
			photo:           user.ProfilePhot,
			birthday:        tmpUserDetail.Born,
			status:          tmpUserDetail.Status,
			job:             tmpUserDetail.Job,
			hometown:        tmpUserDetail.LivesIn,
		}
		userInfos = append(userInfos, userInfo)
	}
	c.JSON(http.StatusOK, gin.H{"data": userInfos})
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
	//var pendingBlogs []Vip_Blog

	//err := Db.Find(&pendingBlogs)
	//if err!=nil || len(pendingBlogs) {
	//	c.JSON(http.StatusOK, gin.H{"pendings": })
	//}
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
