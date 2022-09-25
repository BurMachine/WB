package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type wget struct {
	addr string
}

func (w *wget) GetBody() string {
	req, err := http.Get(w.addr)
	if err != nil {
		log.Fatal("Get error:", err)
	}
	defer req.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Read error:", err)
	}
	return string(b)
}

func main() {
	wgetStruct := new(wget)
	wgetStruct.addr = os.Args[len(os.Args)-1]
	fmt.Println(wgetStruct.GetBody())
}
