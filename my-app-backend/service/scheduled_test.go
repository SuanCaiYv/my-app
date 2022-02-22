package service

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestAddFunc(t *testing.T) {
	params := make(map[string]interface{})
	params["name"] = "test"
	params["age"] = 18
	params["gender"] = "male"
	fn := func() {
		fmt.Println(params["name"])
		fmt.Println(params["age"])
		fmt.Println(params["gender"])
	}
	AddFunc("print", fn)
	AddTask("print", params, time.UnixMilli(time.Now().UnixMilli()+5000))
	params1 := make(map[string]interface{})
	params1["msg"] = "aaa"
	fn1 := func() {
		fmt.Println(params1["msg"])
	}
	AddFunc("print1", fn1)
	AddTask("print1", params1, time.UnixMilli(time.Now().UnixMilli()+10000))
	time.Sleep(math.MaxInt64)
}
