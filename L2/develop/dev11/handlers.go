package main

import (
	"burmachine/calendar/cache/LRU"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res := d.DateFind("day", d.cache.GetAllData(), today)
	for _, re := range res {
		log.Println(*re)
	}
	// Нужно печатать в врайтер!!!!!!!!!!!!!!!!!!
}

func (d *Data) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res := d.DateFind("month", d.cache.GetAllData(), today)
	for _, re := range res {
		log.Println(*re)
	}
	// Нужно печатать в врайтер!!!!!!!!!!!!!!!!!!
}

func (d *Data) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res := d.DateFind("week", d.cache.GetAllData(), today)
	for _, re := range res {
		log.Println(*re)
	}
	// Нужно печатать в врайтер!!!!!!!!!!!!!!!!!!
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

func (d Data) DateFind(str string, data [][]string, today []string) []*jsonStruct {
	res := make([]*jsonStruct, 0)
	for _, d := range data {
		tmp := strings.Split(d[1], "-")
		if DateParse(str, tmp, today) {
			j := &jsonStruct{
				Uuid:  d[0],
				Date:  d[1],
				Event: d[2],
			}
			res = append(res, j)
		}
	}
	return res
}

func DateParse(param string, date, today []string) bool {
	if param == "day" {
		if date[2] == today[2] {
			log.Println(date, "Is fits(day)")
			return true
		} else {
			log.Println(date, "Is not fits(day)")
			return false
		}
	} else if param == "month" {
		if date[1] == today[1] {
			log.Println(date, "Is fits(month)")
			return true
		} else {
			log.Println(date, "Is not fits(month)")
			return false
		}
	} else if param == "week" {
		val1, err := strconv.Atoi(date[2])
		if err != nil {
			fmt.Println("Atoi error in week processing: ", err)
		}
		val2, err := strconv.Atoi(today[2])
		if val1 >= val2 && val1 <= val2+7 {
			log.Println(date, "Is fits(week)")
			return true
		} else {
			log.Println(date, "Is not fits(week)")
			return false
		}
	}
	return true
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
	mux1.HandleFunc("/events_for_week", d.eventsForWeek)
	mux1.HandleFunc("/events_for_month", d.eventsForMonth)

	return mux1
}
