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

	// 数据库
	InitDB()

	// 中间件
	{
		engine.Static("/web/static", "frontend(web)/static")
		engine.Static("/app/static", "frontend(app)/static")
		//		engine.LoadHTMLGlob("frontend(web)/*.html")
	}

	{
		engine.Use(Cors())
		engine.Use(JwtToken())
	}

	// POST请求
	{
		// 登录 提交登录表单
		engine.POST("/login", Login)

		// 注册 提交登录表单
		engine.POST("/register", Register)

		// 上传帖子 (WEB)
		engine.POST("/blog_publish", BlogPublish)

	}

	go func() {
		for {
			msg := <-MsgChan // 改为 Message 类型

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
		webGroup.GET("/profile", ToProfile)

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
		appGroup.GET("/index", ToIndex_app) //blog 中返回了AuthorId

		// POST请求
		{
			// 上传帖子
			appGroup.POST("/upload_post", ToPublish_app)

			// 关注
			appGroup.POST("/follow", ToFollow_app)

			// 喜欢帖子
			appGroup.POST("/like", ToLike_app)
			//评论帖子
			appGroup.POST("/review", ToReview_app)
			// 收藏帖子
			appGroup.POST("/collect", ToCollect_app)
		}

		// 用户相关
		{
			// 获取用户主页
			appGroup.GET("/profile/:uid", ToProfile_app)

			// 个人喜欢列表
			appGroup.GET("/likeList", ToLikeList_app)

			// 个人收藏列表
			appGroup.GET("/collectList", ToCollectList_app)

			// 返回所有朋友
			appGroup.GET("/friends/:uid", ToFriends)

			// 返回所有关注
			appGroup.GET("/follows/:uid", ToFollowing)

			// 返回所有粉丝
			appGroup.GET("/fans/:uid", ToFans)

			// 返回用户排名 (个人)
			appGroup.GET("/rank", ToRank_app)
			appGroup.POST("/upload_info", TouploadInfo_app)
		}

		// 帖子相关
		{
			// 查看帖子详细信息
			appGroup.GET("/post_detail/:Id", func(c *gin.Context) {
				Id := c.Param("Id")
				ToPostDetails_app(Id, c)
			})

			// 首页对帖子标题或内容进行查询
			appGroup.GET("/search/:text", ToSearchedBlogs)
			// 返回朋友帖子
			appGroup.GET("/friend_blog", ToFriendsBlogs)
			//appGroup.GET("/profile_detail")
			appGroup.GET("/show_liked", ToShowLiked_app)
			appGroup.GET("/show_collected", ToShowCollected_app)
			appGroup.GET("/show_reviewed", ToShowReviewed_app)
		}

		// 获取路径信息
		appGroup.GET("/blogPublishTrack", ToPublishTrack_app)

		// 切换到VIP界面，并返回当前VIP的到期时间
		appGroup.GET("/vip", ToVip_app)

	}
	err := engine.Run()
	if err != nil {
		return
	}
}
