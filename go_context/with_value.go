package main

import (
	"context"
	"fmt"
)

func GetValue1(ctx context.Context, key string) {
	value := ctx.Value(key)
	fmt.Println(value)

	ctx1 := context.WithValue(ctx, "key2", "key2")
	GetValue2(ctx1, "key2")
}

func GetValue2(ctx context.Context, key string) {
	value := ctx.Value(key)
	fmt.Println(value)
}

func main() {
	// Trao đổi dữ liệu giữa các func thông qua context value
	// Chỉ sử dụng WithValue cho dữ liệu trong scope request transit API, không phải để truyền tham số tùy chọn cho các hàm.
	ctx := context.WithValue(context.Background(), "key1", "key1")
	GetValue1(ctx, "key1")
}
