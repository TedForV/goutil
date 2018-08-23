package logrushooks

import (
	"fmt"

	"runtime"

	"github.com/sirupsen/logrus"
)

//define additional field in error information
const (
	// ErrorTraceName is the error trace key name
	ErrorTraceName = "Trace"
	// ErrorAddInfoName is the error additional info key name
	ErrorAddInfoName = "AdditionalInfo"
)

// Recover recover the err with params
func Recover(params interface{}) {
	if err := recover(); err != nil {
		trace, n := getTrace()
		logrus.WithField(ErrorAddInfoName, fmt.Sprintf("%+v", params)).WithField(ErrorTraceName, fmt.Sprintf("%s", string((*trace)[:n]))).Error(err)
	}
}

// RecordLog record the log with params
func RecordLog(params interface{}, err error) {
	logrus.WithField(ErrorAddInfoName, fmt.Sprintf("%+v", params)).WithField(ErrorTraceName, fmt.Sprintf("%+v", err)).Error(err)
}

func getTrace() (*[800]byte, int) {
	var trace [800]byte
	n := runtime.Stack(trace[:], false)
	return &trace, n
}
