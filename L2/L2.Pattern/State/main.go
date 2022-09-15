package main

import (
	"log"
	"simple/StatePkg/StatePkg"
)

func main() {
	vendingMachine := StatePkg.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
