package main

import (
	"burmachine/calendar/cache/LRU"
	"burmachine/calendar/configuration"
	"log"
	"net/http"
)

func main() {
	j := new(jsonStruct)
	conf := configuration.New()
	err := conf.LoadConfig("config.yaml")
	if err != nil {
		log.Println("err")
	}
	cache := LRU.NewLRU(100)
	d := &Data{
		json:  *j,
		cache: cache,
	}
	mux := http.NewServeMux()
	mux = d.regMux(mux)
	http.ListenAndServe(conf.Addr, mux)
}
