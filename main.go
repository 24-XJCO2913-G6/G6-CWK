package main

import (
	"github.com/gin-gonic/gin"
	. "main/backend/handlers"
)

func main() {

	engine := gin.Default()

	// 中间件

	// 数据库
	InitDB()

	engine.Static("/web/static", "frontend(web)/static")
	engine.LoadHTMLGlob("frontend(web)/*.html")

	engine.Use(Cors())
	engine.Use(JwtToken())

	engine.POST("/login", Login)
	engine.POST("/register", Register)

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

	appGroup := engine.Group("/app")
	{
		appGroup.GET("/")
		appGroup.GET("/index")
		appGroup.GET("post_detail")
		appGroup.GET("/profile")
		appGroup.GET("/profile_detail")
		appGroup.GET("/rank")

		appGroup.GET("/vip/:uid", func(c *gin.Context) {
			//TODO check whether the user is vip by uid.
			c.JSON(200, gin.H{"isVip": true})
		})
	}

	err := engine.Run()
	if err != nil {
		return
	}
}
