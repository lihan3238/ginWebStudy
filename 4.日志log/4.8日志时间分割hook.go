package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

func (hook FileDateHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02_15-04")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	//有新时间
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)

	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)
	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func InitFile(logPath, appName string) {

	fileDate := time.Now().Format("2006-01-02_15-04")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	filehook := FileDateHook{file, logPath, fileDate, appName}
	logrus.AddHook(&filehook)

}

func main() {
	InitFile("4.日志log/logs_1", "lihan_log")
	for {
		time.Sleep(20 * time.Second)
		logrus.GetLevel()
		logrus.Warnln("warning")

		logrus.Error("error")
		logrus.Infoln("info")
	}

}
