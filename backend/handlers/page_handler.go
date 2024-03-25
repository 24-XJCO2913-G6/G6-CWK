package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ToLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "sign-in.html", gin.H{})
}

func ToRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "sign-up.html", gin.H{})
}

func ToProfile(c *gin.Context, uid string) {

	c.HTML(http.StatusOK, "my-profile.html", gin.H{
		"uid": uid,
	})
}

func ToVip(c *gin.Context) {

	c.HTML(http.StatusOK, "vip.html", gin.H{})
}

func ToBill(c *gin.Context) {

	c.HTML(http.StatusOK, "checkout.html", gin.H{})
}

func ToDownload(c *gin.Context) {

	c.HTML(http.StatusOK, "app-download.html", gin.H{})
}

func ToHelp(c *gin.Context) {

	c.HTML(http.StatusOK, "help.html", gin.H{})
}

func ToHelpDetails(c *gin.Context, searchText string) {

	c.HTML(http.StatusOK, "help-details.html", gin.H{})
}

func ToMessages(c *gin.Context) {

	c.HTML(http.StatusOK, "messaging.html", gin.H{})
}

func ToNotifications(c *gin.Context) {

	c.HTML(http.StatusOK, "notifications.html", gin.H{})
}

func ToSetting(c *gin.Context) {

	c.HTML(http.StatusOK, "settings.html", gin.H{})
}
