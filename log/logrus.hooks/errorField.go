package logrushooks

import (
	"fmt"

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
		// ErrorTraceName is the error trace key name
		logrus.WithField(ErrorAddInfoName, fmt.Sprintf("%+v", params)).WithField(ErrorTraceName, fmt.Sprintf("%+v", err)).Error(err)
	}
}

// RecordLog record the log with params
func RecordLog(params interface{}, err error) {
	// ErrorTraceName is the error trace key name
	logrus.WithField(ErrorAddInfoName, fmt.Sprintf("%+v", params)).WithField(ErrorTraceName, fmt.Sprintf("%+v", err)).Error(err)
}
