package pkg

import (
	"log"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c Card) CheckBalance() error {
	log.Println("Card : Request to Bank")
	time.Sleep(600 * time.Millisecond)
	return c.Bank.CheckBalance(c.Name)
}
