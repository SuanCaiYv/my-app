package util

import (
	"fmt"
	"github.com/google/uuid"
	"mime"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type TaskFunc func(args ...interface{}) (interface{}, error)

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

// UpdateStructObject old必须是指针类型
func UpdateStructObject(old interface{}, m map[string]interface{}) {
	v := reflect.ValueOf(old).Elem()
	t := reflect.TypeOf(old).Elem()
	for i := 0; i < v.NumField(); i += 1 {
		if val, ok := m[t.Field(i).Name]; ok {
			v.Field(i).Set(reflect.ValueOf(val))
		}
	}
}

func UpdateStructObjectWithJsonTag(old interface{}, m map[string]interface{}) {
	v := reflect.ValueOf(old).Elem()
	t := reflect.TypeOf(old).Elem()
	for i := 0; i < v.NumField(); i += 1 {
		if val, ok := m[t.Field(i).Tag.Get("json")]; ok {
			// 处理json的number类型
			if reflect.TypeOf(val).Name() == "float64" && v.Field(i).Type().Name() != "float64" {
				v.Field(i).SetInt(int64(val.(float64)))
			} else {
				v.Field(i).Set(reflect.ValueOf(val))
			}
		}
	}
}

func GetAppPath() string {
	file, err := exec.LookPath(os.Args[0])
	JustPanic(err)
	abs, err := filepath.Abs(file)
	JustPanic(err)
	index := strings.LastIndex(abs, string(os.PathSeparator))
	return abs[:index]
}
