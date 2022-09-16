package StatePkg

import "fmt"

/*
	HasItemState
	Структура состояния, когда идет запрос наличия товара
*/
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (n *HasItemState) AddItem(i int) error {
	fmt.Printf("%d items added\n", i)
	n.vendingMachine.IncrementItemCount(i)
	return nil
}

func (n *HasItemState) RequestItem() error {
	if n.vendingMachine.itemCount == 0 {
		n.vendingMachine.SetState(n.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	n.vendingMachine.SetState(n.vendingMachine.itemRequested)
	return nil
}

func (n *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first!")
}

func (n *HasItemState) DispenseItem() error { // Выдача товара
	return fmt.Errorf("Please select item first!")
}
