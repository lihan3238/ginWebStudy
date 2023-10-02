package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

type MyHook struct {
}

func (hook MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (hook MyHook) Fire(entry *logrus.Entry) error {
	//entry.Data["app"] = "lihan"
	//fmt.Println(entry.Level)

	file, _ := os.OpenFile("4.日志log/error_warn.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	line, _ := entry.String()

	file.Write([]byte(line + "\n"))
	return nil
}

func main() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("warning")
	logrus.Error("error")
	logrus.Debug("debug")
	logrus.Infoln("info")
}
