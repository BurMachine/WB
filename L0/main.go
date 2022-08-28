package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("ui/static/home.page.tmpl")
	if err != nil {
		log.Println("template error:" + err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error() + "\nExecution error")
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Start listening at http://localhost:8080")
	server.ListenAndServe()
}
