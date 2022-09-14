package pkg

import (
	"errors"
	"log"
	"time"
)

/*
	Shop
	Реализовано поведение магазина, взаимодействие с пользователем картой и банком
*/
type Shop struct {
	Name     string
	Products []Product
}

/*
	Sell
	Функция проверки наличия товара и соответствия баланса пользователя
	Является фасадом, который скрывает сложные махинации под собой
*/
func (s Shop) Sell(user User, product Product) error {
	log.Println("Shop: Request to User")
	time.Sleep(500 * time.Millisecond)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	log.Println("Shop : Check", user.Name, "oportunity for buy product...")
	time.Sleep(400 * time.Millisecond)
	for _, p := range s.Products {
		if p != product {
			continue
		}
		if p.Price > user.GetBalance() {
			return errors.New("Shop : No enought money!")
		}
		log.Println("Shop : User", user.Name, "Balance OK")
	}
	return nil
}
