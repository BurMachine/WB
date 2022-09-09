package main

import "fmt"

func BinarySearch(arr []int, elem int) (int, bool) {
	var mid, assumption int

	min := 0 // обозначаем начальный миниму и максимум
	max := len(arr) - 1
	for min <= max { // гоняем цикл, пока не найдем
		mid = (min + max) / 2
		assumption = mid
		if arr[assumption] == elem { // усли этот элемент равен искомому, то возвращаем
			return assumption, true
		}
		if elem > arr[assumption] { // иначе: сравниваем с искомым и рассматриваем ту половину массива,
			min = mid + 1 // в которой предположительно находится нужный нам элемент
		} else {
			max = mid - 1
		}
	}
	return 0, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, ok := BinarySearch(arr, 3)
	fmt.Println(index, ok)
}
