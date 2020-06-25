package main

import (
	"context"
	"fmt"
	"time"
)

// Mong muốn request này complete trong 2s. nếu quá 2s sẽ dừng nó
func RequestAPI() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go func() {
		select {
		// nếu request quá lâu thì channel này sẽ đóng và canceled
		case <-ctx.Done():
			fmt.Println("Request to timeout")
			return
		}
	}()

	// code body xử lý logic
	RequestToServer()

	// check nếu channel close thì dừng function lại
	if err := ctx.Err(); err != nil {
		return err
	}

	fmt.Println("Connect successfully")
	return nil
}

// Giã sử cho nó request mất 2s mới xử lý xong
func RequestToServer() {
	fmt.Println("Connect to Server ....!")
	time.Sleep(time.Second * 2)
}

func main() {
	RequestAPI()
}
