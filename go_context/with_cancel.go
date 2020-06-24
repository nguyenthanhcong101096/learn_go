package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Problem
// what would happen if the client cancelled the request in the middle?
// This could happen if, for example, the client closed their browser mid-request.
// Without cancellation, the application server and database would continue to do their work, even though the result of that work would be wasted:

// Ví dụ hàm Dowload bị fail vì 1 số lí do nào đó
// We use time.Sleep to simulate a resource intensive operation
func DownloadFromGoogle(ctx context.Context) error {
	// set time out khoảng 1 s
	// sau 1s mà chưa complete thì -> error
	timeout, _ := context.WithTimeout(ctx, time.Millisecond*100)

	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	req = req.WithContext(timeout)

	client := &http.Client{}

	res, err := client.Do(req)

	if err == nil {
		fmt.Println("Respone Received status code", res.StatusCode)
		return nil
	}

	fmt.Println("Request fail: ", err)
	time.Sleep(3 * time.Second)

	// giả lập gọi tới google mất hết 100s -> time out
	return err
}

func ExtraFile(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example

	select {
	case <-ctx.Done():
		fmt.Println("Extra cancel by download fail")
	default:
		fmt.Println("Download successfully")
	}
}

func main() {
	// Create a new context, with its cancellation function
	ctx, cancel := context.WithCancel(context.Background())

	// Run two operations: one in a different go routine
	go func() {
		err := DownloadFromGoogle(ctx)

		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	ExtraFile(ctx)
}
