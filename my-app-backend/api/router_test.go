package api

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	fmt.Println(time.Now())
	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			{
				fmt.Println("done")
				fmt.Println(time.Now())
			}
		}
	}()
	wait := make(<-chan struct{})
	<-wait
}
