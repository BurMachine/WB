package main

import "log"

func switchBit(number int64, i, value int) int64 {
	var res int64
	var mask int64 = 1 << i
	if value == 1 {
		res = number | mask
	} else {
		res = number ^ mask
	}
	return res
}

func main() {
	var num int64 = 2
	log.Println(switchBit(num, 0, 1))
}
