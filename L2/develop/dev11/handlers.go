package main

import (
	"burmachine/calendar/cache/LRU"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	json  jsonStruct
	cache *LRU.LRU
}

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

//****************************  POST  ******************************

func (d *Data) createUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("err in create or update", *r)
		return
	}
	d.ParsePostBody(r)

	d.cache.Set(d.json.Uuid, d.json)
}

func (d *Data) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("err in delete", *r)
		return
	}
	d.ParsePostBody(r)
	err := d.cache.Delete(d.json.Uuid)
	if err != nil {
		log.Println("Error in delete handler: ", err)
	}
}

//
//func (d *Data) updateEvent(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		fmt.Println("err in update", *r)
//		return
//	}
//	d.ParsePostBody(r)
//	/*
//		нахожу данные в кеше по ключу и изменяю дату и описание
//		наверное нужно написать функцию в кеше
//		или просто делать сет и он сам будет менять анныее по ключу на более актуальные *
//	*/
//	d.cache.Set(d.json.Uuid, d.json)
//
//}

//****************************  GET  *******************************

func (d *Data) getEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	uid := quer.Get("uid")
	log.Println(d.cache.Get(uid), "вытащил из лру кэша")
	//log.Println(j)
}

func (d *Data) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	d.cache.FindForDaY("day")
}

//*************************************  SUPPORT  ******************************************

func (d *Data) ParsePostBody(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Parse body error: ", err)
		return
	}
	err = json.Unmarshal(body, &d.json)
	if err != nil {
		log.Println("Unmarshalling error in parsing post request's body:\n", err)
		return
	}
}

func (d Data) DateParse() []string {
	res := strings.Split(d.json.Date, "-")
	log.Println(res)
	return res
}

//********************************** REGISTRATION HANDLERS  *********************************S

func (d *Data) regMux(mux *http.ServeMux) *http.ServeMux {
	mux1 := mux
	mux1.HandleFunc("/", helloWorld)
	mux1.HandleFunc("/create_event", d.createUpdateEvent)
	mux1.HandleFunc("/update_event", d.createUpdateEvent)
	mux1.HandleFunc("/delete_event", d.deleteEvent)

	mux1.HandleFunc("/get_event", d.getEvent)
	mux1.HandleFunc("/events_for_day", d.eventsForDay)

	return mux1
}
