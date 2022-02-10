package util

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"path"
	"reflect"
	"testing"
	"time"
)

func TestMIMEType(t *testing.T) {
	fmt.Println(path.Ext("aaa.bbb"))
}

func TestUpdateStructObject(test *testing.T) {
	var result interface{}
	result = &resp.Result{
		Code:      0,
		Msg:       "",
		Data:      nil,
		Timestamp: time.Now(),
	}
	m := map[string]interface{}{"code": 204}
	t := reflect.TypeOf(result).Elem()
	fmt.Println(t.Field(0).Name)
	v := reflect.ValueOf(result).Elem()
	v.Field(0).Set(reflect.ValueOf(m["code"]))
	fmt.Println(result)
}

func TestUpdateStructObjectWithJsonTag(test *testing.T) {
	result := resp.NewOk("aaa")
	m := map[string]interface{}{"code": 204, "msg": "bbb"}
	UpdateStructObjectWithJsonTag(result, m)
	fmt.Println(result)
}
