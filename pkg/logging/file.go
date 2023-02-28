package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogPath() string {
	return fmt.Sprintf("%s%s%s.%s", LogSavePath, LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
}

func mkDir() {
	dir,_ := os.Getwd()
	err := os.MkdirAll(dir + "/" + LogSavePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func openLogFile(filePath string) *os.File {
	_,err := os.Stat(filePath)
	switch {
	case os.IsExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("permission: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to open: %v", err)
	}
	return handle
}