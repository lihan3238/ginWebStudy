package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {

		//c.File("uploads/logo.png")
		//有些响应，比如图片，浏览器就会显示这个图片，而不是下载，所以我们需要使浏览器唤起下载行为
		c.Header("Content-Type", "application/octet-stream")
		//一定要指定下载文件名（可以与源文件名不同），不然默认download无后缀名
		c.Header("Content-Disposition", "attachment; filename=1.png")
		//设置文件传输方式为二进制（乱码问题相关）
		c.Header("Content-Transfer-Encoding", "binary")
		//指定源文件
		c.File("uploads/logo.png")
	})

	router.Run(":8080")
}
