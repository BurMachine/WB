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

//****************************  GET  *******************************

func (d *Data) getEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	uid := quer.Get("uid")
	log.Println(d.cache.Get(uid), "вытащил из лру кэша")
}

func (d *Data) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res, err := d.DateFind("day", d.cache.GetAllData(), today)
	if err != nil {
		http.Error(w, "b logic error", 503)
		return
	}
	err = PrinterJson(res, w)
	if err != nil {
		http.Error(w, "print data error", 500)
		return
	}
}

func (d *Data) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res, err := d.DateFind("month", d.cache.GetAllData(), today)
	if err != nil {
		http.Error(w, "b logic error", 503)
		return
	}
	err = PrinterJson(res, w)
	if err != nil {
		http.Error(w, "print data error", 500)
		return
	}
}

func (d *Data) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("err get method", *r)
		return
	}
	quer := r.URL.Query()
	today := strings.Split(quer.Get("date"), "-")
	res, err := d.DateFind("week", d.cache.GetAllData(), today)
	if err != nil {
		http.Error(w, "b logic error", 503)
		return
	}
	err = PrinterJson(res, w)
	if err != nil {
		http.Error(w, "print data error", 500)
		return
	}
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

func (d Data) DateFind(str string, data [][]string, today []string) ([]*jsonStruct, error) {
	res := make([]*jsonStruct, 0)
	for _, d := range data {
		tmp := strings.Split(d[1], "-")
		ok, err := DateParse(str, tmp, today)
		if err != nil {
			log.Println("Date parse error :", err)
			return nil, err
		}
		if ok {
			j := &jsonStruct{
				Uuid:  d[0],
				Date:  d[1],
				Event: d[2],
			}
			res = append(res, j)
		}
	}
	return res, nil
}

func DateParse(param string, date, today []string) (bool, error) {
	if param == "day" {
		if date[2] == today[2] {
			log.Println(date, "Is fits(day)")
			return true, nil
		} else {
			log.Println(date, "Is not fits(day)")
			return false, nil
		}
	} else if param == "month" {
		if date[1] == today[1] {
			log.Println(date, "Is fits(month)")
			return true, nil
		} else {
			log.Println(date, "Is not fits(month)")
			return false, nil
		}
	} else if param == "week" {
		val1, err := strconv.Atoi(date[2])
		if err != nil {
			fmt.Println("Atoi error in week processing: ", err)
			return false, err
		}
		val2, err := strconv.Atoi(today[2])
		if val1 >= val2 && val1 <= val2+7 {
			log.Println(date, "Is fits(week)")
			return true, nil
		} else {
			log.Println(date, "Is not fits(week)")
			return false, nil
		}
	}
	return true, nil
}

func PrinterJson(res []*jsonStruct, w http.ResponseWriter) error {
	js := new(jsonStruct)
	str := ""
	for _, re := range res {
		log.Println(*re)
		js.Uuid = re.Uuid
		js.Date = re.Date
		js.Event = re.Event
		var b []byte
		b, err := json.Marshal(js)
		if err != nil {
			log.Println("Marshalling error :", err)
			return err
		}
		str += string(b)
		str += "\n"
	}
	_, err := w.Write([]byte(str))
	if err != nil {
		log.Println("Write error :", err)
		return err
	}
	return nil
}

//********************************** MIDDLEWARE  ********************************************

func (d Data) MiddleWareHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL, request.Method)
		next(writer, request)
	})
}

//********************************** REGISTRATION HANDLERS  *********************************

func (d *Data) regMux(mux *http.ServeMux) *http.ServeMux {
	mux1 := mux
	mux1.HandleFunc("/create_event", d.MiddleWareHandler(d.createUpdateEvent))
	mux1.HandleFunc("/update_event", d.MiddleWareHandler(d.createUpdateEvent))
	mux1.HandleFunc("/delete_event", d.MiddleWareHandler(d.deleteEvent))

	mux1.HandleFunc("/get_event", d.MiddleWareHandler(d.getEvent))
	mux1.HandleFunc("/events_for_day", d.MiddleWareHandler(d.eventsForDay))
	mux1.HandleFunc("/events_for_week", d.MiddleWareHandler(d.eventsForWeek))
	mux1.HandleFunc("/events_for_month", d.MiddleWareHandler(d.eventsForMonth))

	return mux1
}
