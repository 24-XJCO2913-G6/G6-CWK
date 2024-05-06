package handlers

import (
	"github.com/gin-gonic/gin"
	. "main/backend/models"
	"net/http"
)

func ToIndex(c *gin.Context) {
	Claim, _ := CheckToken(c.GetString("aToken"))
	Uid := Claim.Uid

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":     c.GetString("message"),
		"aToken":      c.GetString("aToken"),
		"rToken":      c.GetString("rToken"),
		"Uid":         Uid,
		"BlogCount":   BlogCount(c),
		"FollowCount": FollowCount(c),
		"FansCount":   FollowByCount(c),
		"Signature":   SignatureCheck(c),
		"tracks":      RankCheck(c),
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

func ToProfile(c *gin.Context, uid string) {
	c.HTML(http.StatusOK, "my-profile.html", gin.H{
		"uid":     uid,
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
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

func ToManage(c *gin.Context) {
	users := []User{
		{Uid: 1, Name: "Alice", Email: "email1"},
		{Uid: 2, Name: "Bob", Email: "email2"},
		{Uid: 3, Name: "Charlie", Email: "email3"},
	}
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"message": c.GetString("message"),
		"aToken":  c.GetString("aToken"),
		"rToken":  c.GetString("rToken"),
		"users":   users,
	})
}
