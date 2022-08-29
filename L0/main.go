package main

import (
	"fmt"
	"net/http"
	"time"
)

type info struct {
}

func main() {
	dsn := "postgresql://web:123@127.0.0.1:5433/userdb?sslmode=disable"
	_, err := openDB(dsn)
	if err != nil {
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/uid/", handlerView)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Start listening at http://localhost:8080")
	server.ListenAndServe()
}
