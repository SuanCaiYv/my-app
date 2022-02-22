package service

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestAddFunc(t *testing.T) {
	fn1 := func() {
		fmt.Println(time.Now().String())
	}

	AddFunc("print1", fn1)
	curr := time.Now().UnixMilli()
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+5000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+5000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+5000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+6000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+7000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+8000))
	AddTask("print1", make(map[string]interface{}), time.UnixMilli(curr+8000))
	time.Sleep(math.MaxInt64)
}
