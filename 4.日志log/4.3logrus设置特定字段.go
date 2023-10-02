package main

import "github.com/sirupsen/logrus"

func main() {

	log_1 := logrus.WithField("app", "4.3logrus设置特定字段").WithField("service", "logrus")
	log_2 := logrus.WithFields(logrus.Fields{
		"user_id": "22",
		"ip":      "192.168.56.105",
	})
	log_3 := log_1.WithFields(logrus.Fields{
		"user_id": "22",
		"ip":      "192.168.56.105",
	})

	log_1.Errorf("出错了")
	log_2.Errorf("出错了")
	log_3.Errorf("出错了")
}
