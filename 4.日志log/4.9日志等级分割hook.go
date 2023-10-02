package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	alllog  = "all"
	errlog  = "err"
	warnlog = "warn"
	infolog = "info"
)

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warnFile *os.File
	infoFile *os.File
	logPath  string
}

func (hook FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileLevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.errFile.Write([]byte(line))
	case logrus.WarnLevel:
		hook.warnFile.Write([]byte(line))
	case logrus.InfoLevel:
		hook.infoFile.Write([]byte(line))

	}

	hook.file.Write([]byte(line))
	return nil
}

func InitLevel(logPath string) {

	err := os.MkdirAll(fmt.Sprintf("%s", logPath), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	allFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, alllog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	errFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, errlog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	warnFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, warnlog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	infoFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, infolog), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)

	filehook := FileLevelHook{allFile, errFile, warnFile, infoFile, logPath}
	logrus.AddHook(&filehook)

}

func main() {
	InitLevel("4.日志log/log_level")
	logrus.Error("error")
	logrus.Warnln("warning")
	logrus.Infoln("info")
	logrus.Println("print")
}
