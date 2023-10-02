package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func myLogFormat(param gin.LogFormatterParams) string {

	// 你的自定义格式
	return fmt.Sprintf(
		"[lihan]	%s	|%d|	%s%s%s	%s\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.StatusCode,
		//param.Method,
		param.MethodColor(), param.Method, param.ResetColor(), //根据不同的请求类型输出不同颜色
		param.Path,
	)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, numHandlers int) {
		log.Printf(
			"[lihan] %s %s %s %d\n",
			httpMethod,
			absolutePath,
			handlerName,
			numHandlers,
		)
	}
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) //同时写入文件和控制台
	router := gin.New()
	//router.Use(gin.LoggerWithFormatter(myLogFormat))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: myLogFormat}))

	router.GET("/index", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	router.Run(":8080")
}
