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

	engine.GET("/profile/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		ToProfile(c, uid)
	})

	engine.GET("/vip", ToVip)
	engine.GET("/bill", ToBill)

	err := engine.Run()
	if err != nil {
		return
	}
}
