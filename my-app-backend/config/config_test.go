package config

import (
	"fmt"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	obj := ApplicationConfiguration()
	fmt.Println(obj.Roles)
}
