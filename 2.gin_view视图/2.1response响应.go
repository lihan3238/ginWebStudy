package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginString(c *gin.Context) {
	c.String(200, "你好啊！")
	//c.String(http.StatusOK, "你好啊！")
}

func ginJson(c *gin.Context) {
	//json响应结构体
	type UserInfo struct {
		UserName string `json:"username-json"` //返回给前端的字段名
		Age      int    `json:"age_json"`
		PassWord string `json:"-"` //"-"不会返回给前端
	}
	//user := UserInfo{"lihan", 32, "123456"}
	//c.JSON(200, user)

	//json响应map
	//userMap := map[string]string{
	//	"user_name": "lihan",
	//	"age":       "32",
	//}
	//c.JSON(200, userMap)

	//直接响应json
	c.JSON(200, gin.H{"user_name": "lihan", "age": 32})
}

func ginXml(c *gin.Context) {
	c.XML(200, gin.H{"user_name": "lihan", "age": 32, "status": http.StatusOK, "data": gin.H{"id": 1, "name": "lihan"}})
}

func ginYaml(c *gin.Context) {
	c.YAML(200, gin.H{"user_name": "lihan", "age": 32, "status": http.StatusOK, "data": gin.H{"id": 1, "name": "lihan"}})
}

func ginHtml(c *gin.Context) {
	type UserInfo struct {
		UserName string
		Age      int
		PassWord string
	}
	user := UserInfo{"lihan", 32, "123456"}
	c.HTML(200, "index.html", user)

	//c.HTML(200, "index.html", gin.H{"user_name": "lihan", "age": 32, "status": http.StatusOK, "data": gin.H{"id": 1, "name": "lihan"}})
} //gin.H()可以向html传参

func ginRedirect(c *gin.Context) {
	//c.Redirect(302, "/html")
	c.Redirect(301, "https://lihan3238.github.io/")
}

func main() {
	router := gin.Default()
	//加载html模板目录下所有模板文件
	//templates目录要与main.go所在目录同级，而不是在main.go所在目录
	router.LoadHTMLGlob("templates/*")

	//golang中，没有相对文件的路径，只有相对项目的路径

	//配置单个文件，网页请求的路由，文件路径
	router.StaticFile("/downloads/lihan.png", "static/lihan.png")
	//配置文件夹，网页请求的路由，文件夹路径
	router.StaticFS("/downloads/files", http.Dir("static/texts"))
	router.GET("/html", ginHtml)

	router.GET("/string", ginString)
	router.GET("/json", ginJson)
	router.GET("/xml", ginXml)
	router.GET("/yaml", ginYaml)
	router.GET("/lihan", ginRedirect)

	router.Run(":8080")
}
