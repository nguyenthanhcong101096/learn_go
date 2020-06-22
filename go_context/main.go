package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeOut(ctx context.Context) {
	cancelChannel := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			cancelChannel <- true
			return
		}
	}()

	check := <-cancelChannel

	if check {
		close(cancelChannel)
		return
	}

	time.Sleep(time.Second * 1)
	fmt.Println("END")
}

func runWithTimeOut() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	withTimeOut(ctx)
}

// trao đổi dữ liệu giữa các func thông qua context
func withValueA(ctx context.Context) {
	if value := ctx.Value("number"); value != nil {
		ctx := context.WithValue(ctx, "number1", 20)
		withValueB(ctx)
	}
}

func withValueB(ctx context.Context) {
	value1 := ctx.Value("number").(int)
	value2 := ctx.Value("number1").(int)
	fmt.Println(value1 + value2)
}

func runWithValue() {
	ctx := context.WithValue(context.Background(), "number", 10)
	withValueA(ctx)
}

func runWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select {
	case <-ctx.Done():
		fmt.Println("Done rồi nè")
	}
}

func main() {
	runWithCancel()
}
