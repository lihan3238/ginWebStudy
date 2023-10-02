package main

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ArticleInfo struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"lihan", 18},
		{"李寒", 22},
		{"op", 28},
	}
	c.JSON(200, Response{200, userList, "success"})
}

func ArticleListView(c *gin.Context) {
	var articleList []ArticleInfo = []ArticleInfo{
		{"gin", 1},
		{"go", 2},
		{"vue", 3},
	}
	c.JSON(200, Response{200, articleList, "success"})
}

func UserRouterInit(api *gin.RouterGroup) {
	api1 := api.Group("api_1")
	{
		api1.GET("/users1", UserListView)
		api1.POST("/users2", UserListView)
	}
}

func ArticleRouterInit(api *gin.RouterGroup) {
	api2 := api.Group("api_2")
	{
		api2.GET("/users3", ArticleListView)
		api2.POST("/users4", ArticleListView)
	}
}

func main() {
	router := gin.Default()
	api := router.Group("api")

	UserRouterInit(api)
	ArticleRouterInit(api)
	router.GET("/users", UserListView)
	router.Run(":8080")

}
