package filehook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type FileHook struct {
	LogFolderPath string
}

var (
	logFileDay   = 0
	errorLogPath string
	debugLogPath string
)

func NewFileHook(logFolderPath string) *FileHook {
	logFolderPath = strings.Trim(logFolderPath, " ")
	if logFolderPath == "" {
		panic("logFolderPath is null")
	}
	return &FileHook{
		LogFolderPath: logFolderPath,
	}
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	logFilePath := hook.GetLogFilePath(entry.Level)
	fileObj, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer fileObj.Close()

	info, err := entry.String()
	if err != nil {
		fileObj.WriteString(err.Error() + "\r\n")
	} else {
		fileObj.WriteString(info + "\r\n")
	}
	return nil
}

func (hook *FileHook) GetLogFilePath(level logrus.Level) string {
	if logFileDay == 0 || logFileDay != time.Now().Day() { //first time to get path
		defer func() {
			logFileDay = time.Now().Day()
		}()
		errorLogPath = fmt.Sprintf("%s%s", hook.LogFolderPath, getLogName(logrus.ErrorLevel))
		debugLogPath = fmt.Sprintf("%s%s", hook.LogFolderPath, getLogName(logrus.DebugLevel))
	}
	switch level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return errorLogPath
	default:
		return debugLogPath
	}
}

func getLogName(level logrus.Level) string {

	return fmt.Sprintf("%s.%s.log", time.Now().Format("20060102"), getLogFileExtension(level))
}

func getLogFileExtension(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return "error"
	default:
		return "debug"

	}
}