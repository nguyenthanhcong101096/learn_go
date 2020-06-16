### define function

`func func_name (params) return_type { code } `

#### 1. only function no return type

```
func helloNoParams() {
  fmt.Println("Hello world")
}

or with params

func helloWithParams(name string) {
 fmt.Println("hello", name)
}
```

#### 2. function return type

```
func greetString(name string) string {
 result := fmt.Sprintf("Hello %s", name)
 return result
}

Or multiple return values

func multipleReturnValues(w, h int) (int, int, int) {
 area := w * h
 return w, h, area
}

Or named return values

func namedReturnValues(w, h int) (width int, height int, isSquare bool) {
 isSquare = w == h
 return w, h, isSquare
}
```
