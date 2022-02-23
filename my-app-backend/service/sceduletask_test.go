package service

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	AddFunction("print", func(params Params) {
		fmt.Println(params["str"])
	})
	curr := time.Now()

	params1 := make(Params)
	params1["str"] = "hello"
	Add("print", params1, curr.Add(time.Second*2))

	params2 := make(Params)
	params2["str"] = "world"
	Add("print", params2, curr.Add(time.Second*5))

	params3 := make(Params)
	params3["str"] = "foo"
	Add("print", params3, curr.Add(time.Second*3))

	params4 := make(Params)
	params4["str"] = "bar"
	Add("print", params4, curr)

	params5 := make(Params)
	params5["str"] = "msl"
	Add("print", params5, curr.Add(time.Second*6))

	params6 := make(Params)
	params6["str"] = "cwb"
	Add("print", params6, curr.Add(time.Second))

	time.Sleep(math.MaxInt64)
}

func TestAddFunction(t *testing.T) {
	c := make(chan string, 0)
	go func() {
		for {
			select {
			case val := <-c:
				fmt.Println(val)
				time.Sleep(time.Second)
			}
		}
	}()
	go func() {
		c <- "aaa"
	}()
	go func() {
		c <- "bbb"
	}()
	go func() {
		c <- "ccc"
	}()
	go func() {
		c <- "ddd"
	}()
	time.Sleep(math.MaxInt64)
}
