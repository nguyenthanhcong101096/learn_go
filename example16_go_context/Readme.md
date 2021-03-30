## Go Context
Go context có thể tạo bằng 4 cách

- Passing request-scoped values  -  using `WithValue()` function of context package
- With cancellation signals - using `WithCancel()` function of context package
- With deadlines - using `WithDeadine()` function of context package
- With timeouts - using `WithTimeout()` function of context package

## context.WithValue()
Sử dụng để pass request-scoped values.

`withValue(parent Context, key, val interface{}) (ctx Context)`

- Nó nhận vào 1 **context parent** và **key, value** sau đó reture **context** có **{key: value}**. 

```
#Root context
ctxRoot := context.Background()

#Child1 Context chỉ có thể access {"key1": "value1"}
ctxChild1 := context.WithValue(ctxRoot, "key1", "value1")

#Child2 Context có thể access cả key ở ctxChild1 {"key1": "value1", "key2": "value2"}
ctxChild1 := context.WithValue(ctxChild1, "key1", "value1")

```

[Demo Here](https://github.com/nguyenthanhcong101096/training_go/blob/master/go_context/with_value.go)

## context.WithCancel()
Used for cancellation signals

`WithCancel(parent Context) (ctx Context, cancel CancelFunc)`

**context.WithCancel()** sẽ return 2 thứ

- Copy parentContext với 1 cái `done channel` mới
- `cancel function` mà khi được gọi sẽ `close channel` đã thực hiện này

```
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    cancelCtx, cancelFunc := context.WithCancel(ctx)
    go task(cancelCtx)
    time.Sleep(time.Second * 3)
    cancelFunc()
    time.Sleep(time.Second * 1)
}

func task(ctx context.Context) {
    i := 1
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Gracefully exit")
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Println(i)
            time.Sleep(time.Second * 1)
            i++
        }
    }
}
```

```
1
2
3
Gracefully exit
context canceled
```

> task function will gracefully exit once the cancelFunc is called. Once the cancelFunc is called, the error string is set to "context cancelled" by the context package. That is why the output of ctx.Err() is "context cancelled"

[Demo Here](https://github.com/nguyenthanhcong101096/training_go/blob/master/go_context/with_cancel.go)

## context.WithTimeout()
Used for time-based cancellation

`func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)`

**context.WithTimeout()** function will

- Will return a copy of the parentContext with the new done channel.
- Accept a timeout duration after which this done channel will be closed and context will be canceled
- A cancel function which can be called in case the context needs to be canceled before timeout.

```
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)
    defer cancel()
    go task1(cancelCtx)
    time.Sleep(time.Second * 4)
}

func task1(ctx context.Context) {
    i := 1
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Gracefully exit")
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Println(i)
            time.Sleep(time.Second * 1)
            i++
        }
    }
}
```

```
1
2
3
Gracefully exit
context deadline exceeded
```

[Demo Here](https://github.com/nguyenthanhcong101096/training_go/blob/master/go_context/with_timeout.go)

[Nguồn](https://golangbyexample.com/using-context-in-golang-complete-guide/)
