package main

import (
	"github.com/gin-gonic/gin"
	. "main/backend/handlers"
)

func main() {

	engine := gin.Default()

	engine.Static("/static", "frontend/static")
	engine.LoadHTMLGlob("frontend/*.html")

	engine.GET("/", ToIndex)
	engine.GET("/index", ToIndex)

	engine.GET("/login", ToLogin)
	engine.GET("/register", ToRegister)
	engine.POST("/login", Login)
	engine.POST("/register", Register)

	engine.GET("/profile/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		ToProfile(c, uid)
	})

	engine.GET("/vip", ToVip)
	engine.GET("/bill", ToBill)

	engine.GET("/download", ToDownload)

	engine.GET("/help", ToHelp)
	engine.GET("/help/:searchText", func(c *gin.Context) {
		searchText := c.Param("searchText")
		ToHelpDetails(c, searchText)
	})

	engine.GET("/messages", ToMessages)
	engine.GET("/notifications", ToNotifications)
	engine.GET("/setting", ToSetting)

	err := engine.Run()
	if err != nil {
		return
	}
}
