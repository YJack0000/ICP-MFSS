package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/YJack0000/ICP-MFSS/handlers"
	"github.com/YJack0000/ICP-MFSS/middlewares"
)

func main() {
	router := gin.Default()

	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/static/index.html")
	})

	router.POST("/upload", middlewares.RateLimiter(2*time.Second), handlers.UploadFile)
	router.POST("/verify", middlewares.RateLimiter(2*time.Second), handlers.VerifyFile)
	router.GET("/write", handlers.WriteIpToFile)

	router.Run(":8080")
}
