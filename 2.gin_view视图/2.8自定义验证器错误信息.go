package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 获取结构体中的msg参数
func _GetValidMsg(err error, obj any) string {
	// 将err接口断言为具体类型
	//使用的时候传文件指针
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		//循环每一个错误信息
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"用户名校验失败"`
			Age  int    `json:"age" binding:"required,gt=10" msg:"年龄校验失败"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			//c.JSON(200, gin.H{"msg": err.Error()})
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})
	router.Run(":8080")
}
