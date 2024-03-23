package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.Static("/static", "frontend/static")
	engine.LoadHTMLGlob("frontend/*.html")

	err := engine.Run()
	if err != nil {
		return
	}
}
