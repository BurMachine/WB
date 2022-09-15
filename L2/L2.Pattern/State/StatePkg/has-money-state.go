package StatePkg

import "fmt"

/*
	HasMoneyState
	Структура состояния, когда пользователь вносит деньги
*/
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (n *HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (n *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item Dispense in progress.")
}

func (n *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Insert out of stock")
}

func (n *HasMoneyState) DispenseItem() error { // Выдача товара
	fmt.Println("Dispensing item...")
	n.vendingMachine.itemCount -= 1
	if n.vendingMachine.itemCount == 0 {
		n.vendingMachine.SetState(n.vendingMachine.noItem)
	} else {
		n.vendingMachine.SetState(n.vendingMachine.hasItem)
	}
	return nil
}
