package main

import "github.com/sirupsen/logrus"

func main() {

	//行号
	logrus.SetReportCaller(true)
	//设置输出格式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("打印")
}
