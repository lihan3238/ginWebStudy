package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// 单位是字节，<< 是左移预算符号，等价于 8 * 2^20
	// gin对文件上传大小的默认值是32MB
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		readerFile, _ := file.Open()
		writerFile, _ := os.Create("uploads/" + file.Filename)
		defer readerFile.Close()
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)
		fmt.Println(n)
		c.JSON(200, gin.H{"msg": "上传成功！"})
	})

	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("上传成功%d个文件!", len(files))})
	})
	router.Run(":8080")
}
