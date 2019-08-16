package logging

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Level this is Level
type Level int

var (
	// File this is File
	File   *os.File
	logger *logrus.Logger
	// DefaultCallerDepth this is DefaultCallerDepth
	DefaultCallerDepth = 2
	logPrefix          = ""
	filedsData         map[string]interface{}
)

func init() {
	filePath := getLogFileFullPath()
	File = openLogFile(filePath)
	logger = logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
	})
	logger.SetOutput(File)
}

// Debug this is Debug
func Debug(v ...interface{}) {
	filedsData = getLine()
	logger.WithFields(filedsData).Debug(v...)
}

// Info this is Info
func Info(v ...interface{}) {
	filedsData = getLine()
	logger.WithFields(filedsData).Info(v...)
}

// Warn this is Warn
func Warn(v ...interface{}) {
	filedsData = getLine()
	logger.WithFields(filedsData).Warn(v...)
}

// Error this is Error
func Error(v ...interface{}) {
	filedsData = getLine()
	logger.WithFields(filedsData).Error(v...)
}

// Fatal this is Fatal
func Fatal(v ...interface{}) {
	filedsData = getLine()
	logger.WithFields(filedsData).Fatal(v...)
}
func getLine() map[string]interface{} {
	var params = make(map[string]interface{})
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		params["file"] = filepath.Base(file)
		params["line"] = line
	}
	return params
}
