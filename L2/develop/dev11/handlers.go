package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type jsonStruct struct {
	Uuid  string `json:"uuid"`
	Date  string `json:"date"`
	Event string `json:"event"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	str := "Hello World"
	n, e := io.WriteString(w, str)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(n, " ", len(str))
		fmt.Println("Здарова")
	}
}

func (j *jsonStruct) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("err", *r)
		return
	}
	j.ParsePostBody(r)
	log.Println(j)
}

func (j *jsonStruct) ParsePostBody(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Parse body error: ", err)
		return
	}
	err = json.Unmarshal(body, j)
	if err != nil {
		log.Println("Unmarshalling error in parsing post request's body:\n", err)
		return
	}
}

//////

func (j *jsonStruct) regMux(mux *http.ServeMux) *http.ServeMux {
	mux1 := mux
	mux1.HandleFunc("/", helloWorld)
	mux1.HandleFunc("/create_event", j.createEvent)
	return mux1
}
