package util

import (
	"fmt"
	"github.com/google/uuid"
	"mime"
	"path"
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

func VerCode() string {
	return fmt.Sprintf("%d", time.Now().UnixMilli())[:6]
}

func JustPanic(val interface{}) {
	if val != nil {
		logger.Panic(fmt.Sprintf("程序崩溃: %v", val))
	}
}

func MIMEType(filename string) string {
	suffix := path.Ext(filename)
	return mime.TypeByExtension(suffix)
}
