package main

import (
	"html/template"
	"log"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
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

// тут метод гет
func (paket *info) handlerView(w http.ResponseWriter, r *http.Request) {
	myParam := r.URL.Query().Get("UID")
	if text, ok := paket.ma[myParam]; ok {
		//text := paket.ma[myParam]
		w.Write([]byte(text))
	} else {
		log.Println("Нет такой", myParam)
		w.Write([]byte("No such uid exists " + myParam))
	}
}
