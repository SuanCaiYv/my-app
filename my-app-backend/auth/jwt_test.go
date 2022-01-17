package auth

import (
	"fmt"
	"testing"
)

func TestSignRefreshToken(t *testing.T) {
	token, err := SignRefreshToken("2508826394@qq.com")
	fmt.Println(err)
	token, err = SignAccessToken("2508826394@qq.com", "owner")
	fmt.Println(err)
	username, role, err := ValidToken(token)
	fmt.Println(err)
	fmt.Println(username)
	fmt.Println(role)
}
