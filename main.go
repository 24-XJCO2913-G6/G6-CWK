package main

import (
	"github.com/gin-gonic/gin"
	. "main/backend/handlers"
)

func main() {

	engine := gin.Default()

	// 中间件
	//engine.Use(JwtToken())
	//engine.Use(Cors())
	// 数据库
	//InitDB()

	engine.Static("/static", "frontend(web)/static")
	engine.LoadHTMLGlob("frontend(web)/*.html")

	webGroup := engine.Group("/web")
	{
		webGroup.GET("/", ToIndex)
		webGroup.GET("/index", ToIndex)

		webGroup.GET("/login", ToLogin)
		webGroup.GET("/register", ToRegister)
		webGroup.GET("/profile/:uid", func(c *gin.Context) {
			uid := c.Param("uid")
			ToProfile(c, uid)
		})

		webGroup.GET("/vip", ToVip)
		webGroup.GET("/bill", ToBill)

		webGroup.GET("/download", ToDownload)

		webGroup.GET("/help", ToHelp)
		webGroup.GET("/help/:searchText", func(c *gin.Context) {
			searchText := c.Param("searchText")
			ToHelpDetails(c, searchText)
		})

		webGroup.GET("/messages", ToMessages)
		webGroup.GET("/notifications", ToNotifications)
		webGroup.GET("/setting", ToSetting)
	}

	engine.POST("/login", Login)
	engine.POST("/register", Register)

	err := engine.Run()
	if err != nil {
		return
	}
}
