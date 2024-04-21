package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			// 允许 Origin 字段中的域发送请求
			c.Writer.Header().Add("Access-Control-Allow-Origin", c.GetHeader("origin"))
			// 设置预验请求有效期为 70000 秒
			c.Writer.Header().Set("Access-Control-Max-Age", "70000")
			// 设置允许请求的方法
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
			// 设置允许请求的 Header
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, atoken, rtoken")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers, atoken, rtoken")
			// 配置是否可以带认证信息
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// 允许类型校验，不经过网关，放行所有 OPTIONS 请求
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
