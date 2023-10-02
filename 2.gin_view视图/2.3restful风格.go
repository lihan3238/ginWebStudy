package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _bandJson(c *gin.Context, obj any) (err error) {
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

func _getList(c *gin.Context) {
	//包含搜索、分页等功能
	articleList := []ArticleModel{
		{Title: "Go语言入门", Content: "本文是Go语言入门指南"},
		{Title: "Stellarise群星攻略", Content: "本文是Stellarise群星攻略"},
		{Title: "马克思主义学习指南", Content: "本文是马克思注意学习指南"},
		{Title: "李寒个人介绍", Content: "本文是李寒个人介绍"},
	}

	//c.JSON(200, articleList)
	//接口封装
	c.JSON(200, Response{0, articleList, "success"})
}
func _getDetail(c *gin.Context) {
	//获取params中的id
	fmt.Println(c.Param("id"))
	//省略查询数据库的过程
	article := ArticleModel{Title: "李寒个人介绍", Content: "本文是李寒个人介绍"}
	c.JSON(200, Response{0, article, "success"})
}

func _create(c *gin.Context) {
	//接受前端传来的JSON数据
	var article ArticleModel

	err := _bandJson(c, &article)
	if err != nil {
		c.JSON(200, Response{1, nil, "参数错误"})
		return
	}
	//省略插入数据库的过程

	c.JSON(200, Response{0, article, "success"})
}
func _update(c *gin.Context) {
	//获取params中的id
	fmt.Println(c.Param("id"))
	//省略查询数据库的过程
	//接受前端传来的JSON数据
	var article ArticleModel

	err := _bandJson(c, &article)
	if err != nil {
		c.JSON(200, Response{1, nil, "参数错误"})
		return
	}
	//省略插入数据库的过程
	c.JSON(200, Response{0, article, "success"})
}
func _delete(c *gin.Context) {
	//省略查询数据库的过程
	//获取params中的id
	fmt.Println(c.Param("id"))
	//省略删除数据库的过程
	c.JSON(200, Response{0, map[string]string{}, "success"})
}

func main() {
	router := gin.Default()
	router.GET("/articles", _getList)       // 文章列表
	router.GET("/articles/:id", _getDetail) // 文章详情
	router.POST("/articles", _create)       // 新建文章
	router.PUT("/articles/:id", _update)    // 修改文章
	router.DELETE("/articles/:id", _delete) // 删除文章

	router.Run(":8080")
}
