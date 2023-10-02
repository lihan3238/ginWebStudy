package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	//设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(200, gin.H{"data": "看看响应头"})
	})

	router.Run(":8080")
}
