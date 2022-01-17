package db

import (
	"fmt"
	"testing"
)

func TestSystemDaoService_Insert(t *testing.T) {
	instance := NewSysUserDaoService()
	tmp, _ := instance.SelectByUsername("aaa")
	fmt.Println(tmp)
}
