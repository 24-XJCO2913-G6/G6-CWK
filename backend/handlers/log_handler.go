package handlers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetBrowser(c *gin.Context) string {
	userAgent := c.Request.UserAgent()

	if strings.Contains(userAgent, "Firefox") {
		return "Firefox"
	} else if strings.Contains(userAgent, "Chrome") {
		return "Chrome"
	} else if strings.Contains(userAgent, "Safari") {
		return "Safari"
	} else if strings.Contains(userAgent, "Edge") {
		return "Edge"
	} else if strings.Contains(userAgent, "Opera") {
		return "Opera"
	}

	// 其他未知浏览器
	return "Unknown/"
}
