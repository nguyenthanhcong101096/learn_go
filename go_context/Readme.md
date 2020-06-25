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
