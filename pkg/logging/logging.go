package logging

import (
	"fmt"
	"log"
	"note-gin/config"
	"os"

	"runtime"
)

var (
	F             *os.File
	DefaultPrefix = "INFO"
	logPrefix     = ""
	logger        *log.Logger
	levelFlags    = []string{"TRACE", "INFO", "WARN", "ERROR", "FATAL"}
)


type Level int

const (
	TRACE Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func SetUp() {
	filePath := config.Conf.AppConfig.LogFilePath
	F, err := os.Open(filePath)
	if err != nil {
		log.Fatal("logging.Setup err: %v", err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags|log.Lshortfile)
}

func Trace(v ...interface{}) {
	setPrefix(TRACE)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, filename, line, ok := runtime.Caller(0)
	if ok {
		logPrefix = fmt.Sprintf("【%s】-【%s:%d】-", levelFlags[level], filename, line)
	} else {
		logPrefix = fmt.Sprintf("【%s】", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
