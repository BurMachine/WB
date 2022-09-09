package main

import "fmt"

func Quicksort(arr []int) {
	if len(arr) <= 1 { // сортируем до тех пор, пока сортируемая часть массива не будет пустой или состоять из одного элемента
		return
	}

	split := partition(arr)

	Quicksort(arr[:split])
	Quicksort(arr[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2] // берем опорную точку
	left := 0
	right := len(ar) - 1
	for { // проходим массив, чтобы элементы, которые меньше опорной точки оказались слева, в которые больше - справа
		for ; ar[left] < pivot; left++ {
		}
		for ; ar[right] > pivot; right-- {
		}
		if left >= right {
			return right
		}
		swap(ar, left, right)
	}
}
func swap(arr []int, i, j int) { // функция меняет местами два элемента
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := []int{1, 3, 5, 4, 2}
	Quicksort(arr)
	fmt.Println(arr)
}
