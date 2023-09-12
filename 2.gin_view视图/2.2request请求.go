package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 查询参数
func _query(c *gin.Context) {
	fmt.Println(c.Query("username"))
	//c.GetQuery仅判断是否存在，不判断是否为空
	fmt.Println(c.GetQuery("username"))
	//c.QueryArray获取全部username的值，返回一个切片
	fmt.Println(c.QueryArray("username"))
	//c.DefaultQuery获取username的值，如果为空则返回默认值
	fmt.Println(c.DefaultQuery("id", "default_id"))
}

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("username"))
	fmt.Println(c.PostFormArray("id"))
	fmt.Println(c.DefaultPostForm("addr", "default_addr"))
}

// 原始参数
func _rawData(c *gin.Context) {
	//fmt.Println(c.GetRawData())
	body, _ := c.GetRawData()
	fmt.Println(string(body))
}

func bandJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func _rawData2(c *gin.Context) {
	type User struct {
		Username string `json:"name"`
		Password int    `json:"pwd"`
	}
	var user User
	err := bandJson(c, &user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}
func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/rawdata", _rawData)
	router.POST("/rawdata2", _rawData2)

	router.Run(":8080")

}
