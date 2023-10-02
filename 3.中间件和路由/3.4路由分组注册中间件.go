package main

import "github.com/gin-gonic/gin"

type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _UserListView(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var userList []UserInfo = []UserInfo{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 19},
		{Name: "王五", Age: 20},
	}
	c.JSON(200, Res{200, userList, "success"})
}

func Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "3238" {
			c.Next()
			return
		}
		c.JSON(200, Res{401, nil, msg})
		c.Abort()
	}
} //中间件

func _UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	userManager.Use(Middleware("token error")) //user_manager下的所有路由都会经过中间件验证
	{
		userManager.GET("/users", _UserListView) // api/user_manager/users
	}
}

func main() {
	router := gin.Default()

	api := router.Group("api")

	api.GET("/login", func(c *gin.Context) {
		c.JSON(200, Res{200, nil, "login success"})
	}) //不在中间件中的路由，不需要验证

	_UserRouterInit(api)

	router.Run(":8080")

}
