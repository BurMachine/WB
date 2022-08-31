package main

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
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

	// nuts-streaming connection
	sc, err := stan.Connect("test-cluster", "consumer")
	if err != nil {
		fmt.Println(err)
	}
	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		str, _ := separate(m.Data)
		insertDB(string(m.Data), str, db)
		log.Println(str)
	})
	m := mapDB(db)
	fmt.Println(m)
	if err != nil {
		fmt.Println(err)
	}
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Start listening at http://localhost:8080")
	server.ListenAndServe()
}
