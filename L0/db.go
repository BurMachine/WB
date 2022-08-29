package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Не получилось подключиться:\n" + err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("Не получилось пропинговать:\n" + err.Error())
		return nil, err
	}
	return db, nil
}

func existDB(id int, db *sql.DB) bool {

	return true
}
