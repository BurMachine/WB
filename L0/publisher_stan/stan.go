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
	modelList := []string{"../model", "../model1"}
	sc, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		fmt.Println(err)
	}
	defer sc.Close()
	//for i := 0; i < 100; i++ {
	//	sc.Publish("foo", []byte("Здарова"))
	//	time.Sleep(100 * time.Millisecond)
	//}
	mod := new(models.Model)
	for _, s := range modelList {
		file, err := ioutil.ReadFile(s)
		if err != nil {
			log.Println("file is not opened", err)
		}
		json.Unmarshal(file, mod)
	}

}
