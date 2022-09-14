package pkg

/*
	User
	Реализовано поведение пользователя
*/
type User struct {
	Name string
	Card *Card
}

/*
	GetBalance
	Метод для получения баланса для этой карты
*/
func (u *User) GetBalance() float64 {
	return u.Card.Balance
}
