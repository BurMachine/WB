package StatePkg

/*
	VendingMachine
	Структура нашего объекта
*/
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func (v *VendingMachine) AddItem(i int) error {
	return v.currentState.AddItem(i)
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.itemCount = v.itemCount + count
}

/*
	NewVendingMachine
	Создаем объект автомата, задаем необходимые состояния и задаем первоначальное состояние hasItemState.
	То есть - у нас присутствует товар и с ним можно работать.
*/
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}
	itemRequestedState := &ItemRequestState{vendingMachine: v}
	noItemState := &NoItemsState{vendingMachine: v}
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.noItem = noItemState
	v.hasMoney = hasMoneyState
	return v
}
