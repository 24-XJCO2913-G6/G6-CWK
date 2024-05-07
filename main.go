package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "main/backend/handlers"
	. "main/backend/models"
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
	go func() {
		for {
			msg := <-MsgChan // 改为 Message 类型

			// 将消息存储到数据库中
			//insertQuery := "INSERT INTO messages (sender, recipient, message) VALUES (?, ?, ?)"
			//_, err := Db.Exec(insertQuery, msg.Sender, msg.Recipient, msg.Content)
			//if err != nil {
			//	fmt.Printf("Failed to insert message to database: %v", err)
			//	continue
			//}

			// 根据接收人发送消息
			Conns.RLock()
			for Conn, username := range Conns.M {
				if username == msg.Recipient {
					if err := Conn.WriteMessage(websocket.TextMessage, []byte(msg.Content)); err != nil {
						fmt.Printf("Failed to send message: %+v\n", err)
					}
					break
				}
			}
			Conns.RUnlock()
		}
	}()

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
		webGroup.GET("/ws", ToWs)
		webGroup.GET("/post_detail/:Id", func(c *gin.Context) {
			Id := c.Param("Id")
			ToPostDetails(Id, c)
		})
	}

	appGroup := engine.Group("/app")
	{
		appGroup.GET("/")
		appGroup.GET("/index")
		appGroup.GET("/post_detail")
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
