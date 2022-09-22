package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})
	r.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello, " + name,
		})
	})

	r.Run()
}
