package nosql

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Tmp struct {
	Id   string
	Name string
	Age  int
}

func (t *Tmp) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *Tmp) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

func TestRedisClient_Set(t *testing.T) {
	ops := NewRedisClient()
	fmt.Println(ops.Get("tmp"))
}

func TestRedisClient_PushSortQueue(t *testing.T) {
	ops := NewRedisClient()
	tmp1 := Tmp{
		Id:   "aaa",
		Name: "bbb",
		Age:  1,
	}
	ops.PushSortQueue("test", &tmp1, 1)
	tmp2 := Tmp{
		Id:   "ccc",
		Name: "ddd",
		Age:  2,
	}
	ops.PushSortQueue("test", &tmp2, 2)
	fmt.Println(ops.PeeksSortQueue("test"))
}
