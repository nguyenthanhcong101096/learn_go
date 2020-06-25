package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Problem
// what would happen if the client cancelled the request in the middle?
// This could happen if, for example, the client closed their browser mid-request.
// Without cancellation, the application server and database would continue to do their work, even though the result of that work would be wasted:

// hàm download file về. giã sử vì 1 số lí do gì bị download fail
func Download(ctx context.Context, file chan int) bool {
	done := 0
	for {
		done++
		file <- done
		red := color.New(color.FgGreen)
		red.Print("+")

		if done == 50 {
			return false
		}
	}

	return false
}

func Compress(ctx context.Context, file chan int) {
	for {
		select {
		case <-file:
			fmt.Print("-")
			time.Sleep(time.Millisecond * 200)
		case <-ctx.Done():
			fmt.Println(" EORR : request has cancel")
			close(file)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	file := make(chan int)

	fmt.Println("Download and Compress")

	// Run two operations: one in a different go routine
	go func() {
		ok := Download(ctx, file)

		// If this operation returns an error
		// cancel all operations using this context
		if !ok {
			cancel()
		}
	}()

	Compress(ctx, file)
}
