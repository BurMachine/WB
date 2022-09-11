package handlers

import "fmt"

// Есть интерфейс, под который нам нужно "адаптироваться"
// В нем есть метод, нам нужно будет реализовать, чтобы соответствовать этому интерфейсу
type Dater interface {
	GetData()
}

// DataCollector - its
type DataCollector struct {
	Num int
}

func (d *DataCollector) GetData() {
	fmt.Println(d.Num)
}
