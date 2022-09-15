package StatePkg

import "fmt"

/*
	ItemRequestState
	Структура состояния, которое говорит о том, что товар уже был запрошен
*/
type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (n *ItemRequestState) AddItem(i int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (n *ItemRequestState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (n *ItemRequestState) InsertMoney(money int) error {
	if money < n.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert : ", n.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is OK")
	n.vendingMachine.SetState(n.vendingMachine.hasMoney)
	return nil
}

func (n *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("Please insert money first!")
}
