package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//获取请求头
	router.GET("/", func(c *gin.Context) {
		//单词首字母大小写不区分，单词之间用"-"分割
		//用于获取一个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.Request.Header.Get("User-Agent"))
		//fmt.Println(c.GetHeader("user-agent"))
		//fmt.Println(c.GetHeader("user-Agent"))

		//Header是一个map[string][]string类型
		fmt.Println(c.Request.Header)
		//获取所有请求头,区分大小写
		fmt.Println(c.Request.Header["User-Agent"])

		c.JSON(200, gin.H{"msg": "ok"})
	})

	//利用请求头，将爬虫和用户区别对待
	//
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		//方法一 正则去匹配
		//字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			//爬虫来了
			c.JSON(200, gin.H{"data": "这是一个爬虫"})
			return

		}
		c.JSON(200, gin.H{"data": "这是一个用户"})
	})
	router.Run(":8080")
}
