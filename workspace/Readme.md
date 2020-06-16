- set $GOPATH

```
$ pwd
$ /Users/congnt/Desktop/golang/workspace
$ export GOPATH=/Users/congnt/Desktop/golang/workspace
$ echo $GOPATH
```

Cấu trúc:

```
workspace
|
|----- src
|      |
|      |-- project_name
|      |        |
|      |        |-- main.go      // go install -> create bin/project_name
|      |
|      |-- validator
|               |-- validator.go // go install -> create pkg/darwin_amd64/validator.a
|      
|---- bin
|      |-- project_name(file)
|
|---- pkg
|      |
|      |-- darwin_amd64
|               |-- validator.a
|     
```

- Trong project chỉ có 1 file `main.go`
- bin/ đống gói toàn src
- pkg/ đống gói function -> import to main.go
