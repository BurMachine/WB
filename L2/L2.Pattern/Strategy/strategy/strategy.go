package strategy

/*
	Strategy
	Поведенческий паттерн, позволяющий выбор поведения алгоритма в ходе исполнения.
	Этот паттерн определяет алгоритмы,инкапсулирует их и использует их взаимозаменяемо
*/

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator
}

func (o Operation) Operate(left, right int) int {
	return o.Operator.Apply(left, right)
}

// Addiction
// Сложение
type Addiction struct {
}

func (Addiction) Apply(lVal, rVal int) int {
	return lVal + rVal
}

// Multiplication
// Умножение
type Multiplication struct {
}

func (Multiplication) Apply(lVal, rVal int) int {
	return lVal * rVal
}
