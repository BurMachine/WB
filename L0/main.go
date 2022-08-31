package main

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/stan.go"
	"net/http"
	"time"
)

type info struct {
	db *sql.DB
	ma map[string]string
}

func main() {
	dsn := "postgresql://web:123@127.0.0.1:5433/userdb?sslmode=disable"
	db, err := openDB(dsn)
	if err != nil {
		return
	}
	ma := mapDB(db)
	a := &info{
		db: db,
		ma: ma,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/uid/", a.handlerView)
	// кеш начальный

	// nuts-streaming connection
	sc, err := stan.Connect("test-cluster", "consumer")
	if err != nil {
		fmt.Println(err)
	}
	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		str, _ := separate(m.Data)
		insertDB(string(m.Data), str, db)

		ma = mapDB(db)
	})
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
