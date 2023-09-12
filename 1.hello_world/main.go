package main

import (
	"github.com/gin-gonic/gin"
)

func Index2(context *gin.Context) {
	context.String(200, "Hello lihan2!")
}

func main() {

	//创建一个默认路由引擎
	router := gin.Default()

	//注册一个路由和处理函数，访问/index的路由时，会执行后面的匿名函数
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "Hello lihan!")
	})

	//另一种方法，可以直接使用已经声明的函数
	router.GET("/index2", Index2)

	//启动HTTP服务,gin会默认把web服务器运行在本机的0.0.0.0:8080端口上(即所有网卡IP的8080端口)
	router.Run(":8080")
	//第二种启动方式，用原生http服务的方式启动，这种方式可以实现更多的自定义配置
	//http.ListenAndServe(":8080", router)
}
