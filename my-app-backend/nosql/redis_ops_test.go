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
}

func TestRedisClient_PushSortQueue(t *testing.T) {
	redis := NewRedisClient()
	//t1 := Tmp{
	//	Id:   "aaa",
	//	Name: "aaa",
	//	Age:  1,
	//}
	//redis.PushSortQueue("test", &t1, 1)
	params := make(map[string]interface{})
	params["aaa"] = "aaa"
	params["bbb"] = 2
	err := redis.PushSortQueue("test", params, 1)
	fmt.Println(err)
}
