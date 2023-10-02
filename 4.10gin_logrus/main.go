package main

import (
	"gowebstudy/4.10gin_logrus/log"
	"gowebstudy/4.10gin_logrus/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log.InitFile("4.10gin_logrus/log", "gin_log")
	router := gin.New()
	router.Use(middleware.LogMiddleware())
	router.GET("/", func(c *gin.Context) {
		logrus.Info("coming!")
		c.JSON(200, gin.H{"message": "ok"})
	})

	router.Run(":8080")
}
