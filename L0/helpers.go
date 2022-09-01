package main

import (
	"encoding/json"
	"github.com/BurMachine/WB/L0/models"
	"log"
)

func separate(str []byte) (string, error) {
	mod := new(models.Model)
	err := json.Unmarshal(str, mod)
	if err != nil {
		log.Println("json unmarshalling error int helper function", err)
		return "", err
	}
	return mod.Order_uid, nil
}
