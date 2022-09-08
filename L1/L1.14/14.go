package main

import "fmt"

func TypeDetermination(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	case chan interface{}:
		fmt.Println("chan type")
	}
}

func main() {
	var value interface{}
	value = true
	//value = 12
	//value = "qwerty"
	//value := make(chan interface{})
	TypeDetermination(value)
}
