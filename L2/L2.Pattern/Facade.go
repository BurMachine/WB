package main

import (
	"log"
	"simple/pkg/pkg"
)

/*
	Это структурный паттерн проектирования, который упрощает доступ к сложной системе
	Фасад инкапсулирует сложную подсистему в один (более простой интерфейсный объект)
*/
var (
	bank = pkg.Bank{
		Name:  "Bank123456",
		Cards: []pkg.Card{},
	}
	Card1 = pkg.Card{
		Name:    "crd-1",
		Balance: 1000,
		Bank:    &bank,
	}
	Card2 = pkg.Card{
		Name:    "crd-2",
		Balance: 2000,
		Bank:    &bank,
	}
	user1 = pkg.User{
		Name: "user1",
		Card: &Card1,
	}
	user2 = pkg.User{
		Name: "user2",
		Card: &Card2,
	}
	prod = pkg.Product{
		Name:  "Cheese",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "24hour",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	bank.Cards = append(bank.Cards, Card1, Card2) // Добавление банком карт
	log.Println("Bank created some cards!")
	log.Println("User: ", user1.Name)
	err := shop.Sell(user1, prod) // Фасад
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("User: ", user2.Name)
	err = shop.Sell(user2, prod) // Фасад
	if err != nil {
		log.Println(err)
		return
	}
}
