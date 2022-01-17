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
	tmp := Tmp{}
	ops := NewRedisClient()
	ops.Get("tmp", &tmp)
	fmt.Println(tmp)
}
