package util

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

var logger = NewLogger()

func GenerateUUID() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		logger.Error("生成UUID失败")
		return time.Now().String()
	}
	return newUUID.String()
}

func JustPanic(val interface{}) {
	if val != nil {
		logger.Panic(fmt.Sprintf("程序崩溃: %v", val))
	}
}
