package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	//自定义验证器 sign
	Name string `form:"name" binding:"required,sign" msg:"请输入名字"`
	Age  int    `form:"age" binding:"required,gt=10" msg:"请输入年龄"`
}

func signValid(fl validator.FieldLevel) bool {
	//不允许name为nameList中的值
	var nameList []string = []string{"lihan", "lihan3238", "Lihan"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true

}

func main() {
	router := gin.Default()
	//自定义验证器 sign
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
	router.POST("/", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(200, err.Error())
			return
		}
		c.JSON(200, gin.H{"data": user})
		return
	})

	router.Run(":8080")
}
