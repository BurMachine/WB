package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurMachine/WB/L0/models"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

func main() {
	modelList := []string{"../model.json", "../model1.json", "../model3.json", "../model_err.json"}
	sc, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		fmt.Println(err)
	}
	defer sc.Close()
	mod := new(models.Model)
	sliceMod := make([]models.Model, 0)
	for _, s := range modelList {
		file, err := ioutil.ReadFile(s)
		if err != nil {
			log.Println("file is not opened", err)
		}
		err = json.Unmarshal(file, mod)
		if err != nil {
			log.Println("Publisher: input error", err)
			continue
		}
		sliceMod = append(sliceMod, *mod)
		err = sc.Publish("foo", file)
		if err != nil {
			log.Println("Publisher: sc.Publish() function error\n", err)
		}
	}
}
