package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type info struct {
	db *sql.DB
}

func main() {
	dsn := "postgresql://web:123@127.0.0.1:5433/userdb?sslmode=disable"
	db, err := openDB(dsn)
	if err != nil {
		return
	}
	a := &info{
		db: db,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/uid/", a.handlerView)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Start listening at http://localhost:8080")
	server.ListenAndServe()
}
