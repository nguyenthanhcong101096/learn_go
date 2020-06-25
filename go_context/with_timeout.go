package main

import (
	"context"
	"fmt"
	"time"
)

// Muốn request này complete trong 2s, nếu sau 2s thì sẽ tự cancel
func RequestAPI() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*6)
	wait := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			wait <- true
			break
		}
	}()

	<-wait

	//giã sử cho nó request mấy 3s mới xử lý xong
	time.Sleep(time.Second * 3)
}

func main() {
	RequestAPI()
}
