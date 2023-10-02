package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Age  int
}

func m10(c *gin.Context) {
	fmt.Println("m10......in")
	c.Set("name", "lihan")
	c.Set("user", Person{
		Name: "lihan",
		Age:  18,
	})
}

func main() {
	router := gin.Default()
	router.Use(m10) // 全局注册中间件

	router.GET("/", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)

		_user, _ := c.Get("user")
		user, ok := _user.(Person) // 使用类型断言
		if !ok {
			fmt.Println("类型断言失败")
		} else {
			fmt.Println(user.Name, user.Age)
		}

		c.JSON(200, gin.H{"msg": "index1"})

	})
	router.Run(":8080")
}
