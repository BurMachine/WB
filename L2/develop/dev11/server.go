package main

import (
	"burmachine/calendar/cache/LRU"
	"net/http"
)

func main() {
	j := new(jsonStruct)
	cache := LRU.NewLRU(100)
	d := &Data{
		json:  *j,
		cache: cache,
	}
	mux := http.NewServeMux()
	mux = d.regMux(mux)
	http.ListenAndServe(":8080", mux)
}
