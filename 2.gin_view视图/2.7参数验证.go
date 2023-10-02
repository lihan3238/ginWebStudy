package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	// binding:"required"不能为空或不传
	Name string `json:"name" binding:"required"` //用户名
	Mail string `json:"mail"`                    //邮箱
	// binding:"min=6,max=12"最小长度6，最大长度12
	Password   string   `json:"password" binding:"min=6,max=12"`                     //密码
	RePassword string   `json:"re_password" binding:"eqfield=Password"`              //确认密码
	Sex        string   `json:"sex" binding:"oneof=男 女"`                             //性别
	HobbyList  []string `json:"hobby_list" binding:"required,dive,startswith=ilove"` //爱好
	IP         string   `json:"ip" binding:"ip"`                                     //ip地址
	//必须用datetime=2006-01-02 15:04:05这个时间，不能换成别的时间
	//1月2号下午3点4分5秒2006年
	Date string `json:"date" binding:"datetime=2006-01-02 15:04:05"` //日期
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, err.Error())
			return
		}

		c.ShouldBindJSON(&user)
		c.JSON(200, user)
	})

	router.Run(":8080")
}
