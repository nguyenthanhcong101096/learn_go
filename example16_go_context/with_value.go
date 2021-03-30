package main

import (
	"context"
	"net/http"
)

func Welcome(rw http.ResponseWriter, r *http.Request) {
	//handler này sẽ lấy value từ context middlware
	//pass request-scope value MiddleWareCheckLogin

	user := r.Context().Value("user_1")
	if user != nil {
		rw.Write([]byte("Hello sir"))
	} else {
		rw.Write([]byte("Hello guide"))
	}
}

func MiddleWareCheckLogin(next http.Handler) http.Handler {
	check := func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), "user_1", "ttlcong")
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	}

	return http.HandlerFunc(check)
}

func main() {
	welcomeHalder := http.HandlerFunc(Welcome)
	http.Handle("/welcome", MiddleWareCheckLogin(welcomeHalder))
	http.ListenAndServe(":8080", nil)
}
