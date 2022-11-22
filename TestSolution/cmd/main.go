package main

import (
	"burmachine/TestSolution/internal/config"
	"burmachine/TestSolution/internal/handlers"
	myStorage "burmachine/TestSolution/storage"
	"flag"
	_ "github.com/gorilla/mux"
	gorillaMux "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	if *cfgPath == "" {
		log.Fatalln("Path to configuration file was not provided")
	}

	cfg, err := config.NewConfigurationFromFile(cfgPath)
	if err != nil {
		log.Fatalln("Configuration getting error")
	}

	storage := myStorage.NewStorage()
	handlerStorage := new(handlers.Storage)
	handlerStorage.StorageH = storage

	muxSet := gorillaMux.NewRouter()
	muxGet := gorillaMux.NewRouter()

	muxSet.HandleFunc("/", handlerStorage.SetHandler).Methods("POST")
	muxGet.HandleFunc("/", handlers.GetHandler).Methods("GET")

	go func() {
		err = http.ListenAndServe(cfg.Port2, muxGet)
		if err != nil {
			log.Fatalln("error with server")
		}
	}()
	err = http.ListenAndServe(cfg.Port1, muxSet)
	if err != nil {
		log.Fatalln("error with server")
	}
}
