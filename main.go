package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("src/*.html")

	router.Static("/assets/", "src/")

	router.GET("/", authHandler)

	_ = router.Run("127.0.0.1:8080")
}

func authHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
