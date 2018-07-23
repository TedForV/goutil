package logrushooks

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//define additional field in error information
const (
	ERROR_TRACE_NAME   = "Trace"
	ERROR_ADDINFO_NAME = "AdditionalInfo"
)

func Recover(params interface{}) {
	if err := recover(); err != nil {
		logrus.WithField(ERROR_ADDINFO_NAME, params).WithField(ERROR_TRACE_NAME, fmt.Sprintf("%+v", err)).Error(err)
	}
}
