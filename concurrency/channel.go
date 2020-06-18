package main

import "fmt"

// channel
// 1. Unbuffered channel
// 2. Buffered channel
// 3. select
// 4. close channel

func unbufferedChannel() {
	ch := make(chan int)

	go func() {
		ch <- 100 // bi block ở đây nếu k có lấy giá trị nó
	}() // gửi giá trị về channel

	fmt.Println(<-ch) // lấy giá trị từ channel // k có giá nào đẫy vào sẽ bị đeadlock here
	fmt.Println("Done")
}

func bufferedChannel() {
	ch := make(chan int, 2) // đăng kí 2 giá trị vào channel này

	ch <- 1
	ch <- 2

	close(ch) // close thì khi lấy value từ channel thì giá trị bằng k / 1-2-0

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sellectChannel() {
	// được sử dụng ở chương trình khi tao có nhiều goroutines và quan tâm đến giá trị trả về

	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true
	}()

	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func main() {
	sellectChannel()
}
