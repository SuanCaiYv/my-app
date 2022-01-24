package main

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/api"
	"time"
)

func main() {
	fmt.Println(int64(time.Second / time.Millisecond))
	BeforeStart()
	api.Route()
}
