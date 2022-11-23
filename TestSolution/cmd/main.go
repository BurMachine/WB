package main

import (
	"burmachine/TestSolution/internal/auth"
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
	usersPath := flag.String("users", "./users.yaml", "Path to yaml users file")
	if *cfgPath == "" {
		log.Fatalln("Path to configuration file was not provided")
	}

	cfg, err := config.NewConfigurationFromFile(cfgPath)
	if err != nil {
		log.Fatalln("Configuration getting error")
	}

	usr, err := auth.NewAuthData(usersPath)
	if err != nil {
		log.Fatalln("Path to users file was not provided")
	}

	authMiddleware := new(handlers.AuthMiddlwareData)
	authMiddleware.Data = usr

	storage := myStorage.NewStorage()
	handlerStorage := new(handlers.Storage)
	handlerStorage.StorageH = storage

	muxSet := gorillaMux.NewRouter()
	muxGet := gorillaMux.NewRouter()

	muxSet.HandleFunc("/", authMiddleware.BasicAuth(handlerStorage.SetHandler)).Methods("POST")
	muxGet.HandleFunc("/", handlerStorage.GetHandler).Methods("GET")

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
