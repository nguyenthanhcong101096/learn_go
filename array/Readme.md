### ARRAY - Cấu trúc dữ liệu mảng

#### Khái báo array

```
var myArray [3]int
```

#### khai báo mảng không cần size
```
myArray := [...]string {"IPHONE", "SAMSUNG"}
```

#### Gán giá trị cho array
```
myArray[0] = 1
```

#### Khai báo array có gía trị khởi tạo
```
myArray := [2]int {1,3}
```

#### Size của array

```
len(myArray)
```

#### Loop array
```
for i := 0; i < len(myArray); i++ {
	fmt.Println(myArray[i])
}

OR

for index, _ := range myArray {
  fmt.Println(myArray[index])
}

OR

for index, value := range myArray {
  fmt.Println(myArray[index])
}
```

#### Mang 2 chiều
```
matrix := [4][2]int {
  {1, 2},
  {3, 4},
  {5, 6},
  {7, 8},
}
```