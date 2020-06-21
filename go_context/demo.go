package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*10)

	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Request fail", err)
		return
	}

	fmt.Println("Respone Received status code", res.StatusCode)
}
