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
	//stmt := `INSERT INTO snippets (title, content_, created, expires)
	//	VALUES($1, $2, current_timestamp, current_timestamp + interval '1 year' * $3)`
	//paket.db.Exec(stmt, "qwerty", "abc", "123")
	myParam := r.URL.Query().Get("UID")
	if a, _ := existDB(myParam, paket.db); a {
		_, text := selectJSON(myParam, paket.db)

		w.Write([]byte(text))
	}
	w.Write([]byte("You enter" + myParam))
}
