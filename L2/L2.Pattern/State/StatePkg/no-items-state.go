package StatePkg

import "fmt"

/*
	NoItemsState
	Структура состояния, когда в автомате отсутствуют товары.
	Оно инкапсулирует вендинговую машину, чтоб мы могли с ней работать
*/
type NoItemsState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemsState) AddItem(i int) error {
	n.vendingMachine.IncrementItemCount(i)              // Увеличиваем кол-во товара
	n.vendingMachine.SetState(n.vendingMachine.hasItem) // Ставим флажок, что товар есть в наличии
	return nil
}

func (n *NoItemsState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemsState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemsState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
