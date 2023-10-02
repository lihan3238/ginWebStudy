package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "m1"})
	fmt.Println("m1...in")
	//next前是请求
	c.Next()
	//next后是响应
	fmt.Println("m1...out")
}

func main() {
	router := gin.Default()

	router.GET("/", m1, func(c *gin.Context) {

		c.JSON(200, gin.H{"msg": "index"})
		fmt.Println("index...in")
		c.Next()
		fmt.Println("index...out")
	})

	router.Run()
}
