package main

import (
	"fmt"
	"log"
	"strconv"
)

func unpacking(str string) []string {
	runes := []rune(str)
	res := make([]string, 0)
	for i := 0; i < len(runes); i++ {
		if runes[i] > '0' && runes[i] < '9' {
			if i != 0 {
				num, err := strconv.Atoi(string(runes[i]))
				if err != nil {
					fmt.Errorf("Atoi error", err)
					return res
				}
				for j := 0; j < num-1; j++ {
					res = append(res, string(runes[i-1]))
				}
			} else {
				log.Println("INCORRECT STRING")
				return res
			}
			continue
		} else if runes[i] == 92 && i < len(runes) {
			log.Println("Здарова")
		}
		res = append(res, string(runes[i]))
	}
	return res
}

func main() {
	//str := "a2b3cd"
	//str := "abcd"
	//str := "45"
	//str := ""
	str := "qwe\\5"
	log.Println(unpacking(str))
}
