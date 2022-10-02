package main

import (
	"log"
	"strings"
)

func main() {
	str := "2022-10-11"
	res := strings.Split(str, "-")
	log.Println(res)
}
