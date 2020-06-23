package main

import (
	"context"
	"log"
	handlers "microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handler
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	router := mux.NewRouter()

	router.HandleFunc("/", ph.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/product", ph.AddProducts).Methods(http.MethodPost)
	router.HandleFunc("/product/{id:[0-9]+}", ph.UpdateProducts).Methods(http.MethodPut)

	// create new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Receiver", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
