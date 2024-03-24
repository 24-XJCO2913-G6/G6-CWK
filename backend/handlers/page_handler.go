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
