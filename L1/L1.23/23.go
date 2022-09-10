package main

import "fmt"

// Удление элемента с соранением порядка
func deleteOrder(index int, slice []int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Удаление элемента без сохранения порядка
func deleteWithoutOrder(index int, slice []int) []int {
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	arr := []int{1, 2, 3, 4, 777, 5}
	//arr = deleteWithoutOrder(4, arr)
	arr = deleteOrder(4, arr)
	fmt.Println(arr)
}
