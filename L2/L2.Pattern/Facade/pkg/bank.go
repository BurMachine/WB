package pkg

import (
	"errors"
	"log"
	"time"
)

/*
	Bank
	Реализовано поведение банка
*/
type Bank struct {
	Name  string
	Cards []Card
}

// CheckBalance
//Проверка баланса на карте
func (b Bank) CheckBalance(CardNumber string) error {
	log.Println("Bank : Request to balance " + CardNumber)
	time.Sleep(300 * time.Millisecond)
	for _, card := range b.Cards {
		if card.Name != CardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("Bank : Balance < 0")
		}
	}
	log.Println("Bank : Balance > 0")
	return nil
}
