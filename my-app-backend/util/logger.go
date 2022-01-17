package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func NewLogger() *logrus.Logger {
	formatter := &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                true,
		DisableQuote:              false,
		EnvironmentOverrideColors: true,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05.000000",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
		},
	}
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	return logger
}
