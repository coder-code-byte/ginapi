package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// LogSavePath this is LogSavePath
	LogSavePath = "/var/run/ginlogs/"
	// LogSaveName this is LogSaveName
	LogSaveName = "log"
	// LogFileExt this is LogFileExt
	LogFileExt = "log"
	// TimeFormat this is TimeFormat
	TimeFormat = "2006-01-02"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	// dir, _ := os.Getwd()
	err := os.MkdirAll(getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
