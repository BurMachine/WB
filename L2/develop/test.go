package main

import "time"
import "log"
import "strings"

func main()  {
	dt := time.Now()
	tmp:= strings.Split(dt.String()," ")
	tmp1 := strings.Split(tmp[0], "-")
	log.Println(tmp1)
}