package main

import (
	"net/http"
)

func main() {
	j := new(jsonStruct)
	mux := http.NewServeMux()
	mux = j.regMux(mux)
	http.ListenAndServe(":8080", mux)
}
