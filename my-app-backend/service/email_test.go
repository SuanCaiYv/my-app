package service

import (
	"testing"
)

func TestSendVerCode(t *testing.T) {
	SendVerCode("2508826394@qq.com", "123456")
}
