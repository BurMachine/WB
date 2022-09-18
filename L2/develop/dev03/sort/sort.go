package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Info struct {
	k        int
	n        bool
	r        bool
	u        bool
	fileName string
}

func (flags Info) openFile() *os.File {
	file, err := os.Open(flags.fileName)
	if err != nil {
		log.Fatal("File not open: ", err)
	}
	return file
}

func (flags *Info) fileProc(file *os.File) []string {
	fileScanner := bufio.NewScanner(file)
	res := make([]string, 1)
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		res = append(res, fileScanner.Text())
	}
	return res
}

func main() {
	k := flag.Int("k", -1, "Column to sort")
	n := flag.Bool("n", false, "Number sort")
	r := flag.Bool("r", false, "Reverse sorting")
	u := flag.Bool("u", false, "Without duplicate")
	flag.Parse()
	info := &Info{
		k:        *k,
		n:        *n,
		r:        *r,
		u:        *u,
		fileName: os.Args[len(os.Args)-1],
	}

	file := info.openFile()
	res := info.fileProc(file)
	fmt.Println(res)
	info.Sorting(&res)
	for _, re := range res {
		if re != "" {
			fmt.Println(re)
		}
	}
}

func (flags *Info) Sorting(slice *[]string) {

	if flags.r {
		sort.Strings(*slice)
		RSorting(*slice)
	}

	if flags.u {
		sort.Strings(*slice)
		*slice = USorting(*slice)
	}

	if flags.k > -1 {
		*slice = KSorting(*slice, flags.k)
	}

	if flags.n {
		return
	}
}

func RSorting(text []string) {
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]
	}
}

func USorting(text []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0)
	a := struct{}{}
	for _, s := range text {
		_, ok := m[s]
		if !ok {
			m[s] = a
			res = append(res, s)
		}
	}
	return res
}

func KSorting(text []string, k int) []string {
	tmp := make([]string, 0)
	res := make([]string, 0)
	tmp1 := ""
	tmp2 := make([]string, 0)
	/*
		Сплит строки
		Замена слова[k] со словом[0]
		Соеденить обратно воедино
		Отсортировать
		Разделить
		Поменять все на место
		Соеденить обратно
	*/
	for j, s := range text {
		if j == 0 {
			continue
		}
		tmp = append(tmp, strings.Split(s, " ")...)
		tmp[k], tmp[0] = tmp[0], tmp[k]
		for i, val := range tmp {
			tmp1 += val
			if i != len(tmp) {
				tmp1 += " "
			}
		}
		tmp2 = append(tmp2, tmp1)
		tmp = tmp[:0]
		tmp1 = ""
	}
	sort.Strings(tmp2)
	// Возвращаю все на место

	for _, s := range tmp2 {
		tmp = append(tmp, strings.Split(s, " ")...)
		tmp[k], tmp[0] = tmp[0], tmp[k]
		for i, val := range tmp {
			tmp1 += val
			if i != len(tmp) {
				tmp1 += " "
			}
		}
		res = append(res, tmp1)
		tmp = tmp[:0]
		tmp1 = ""
	}
	//for i := 0; i < len(res); i++ {
	//	fmt.Println(res[i])
	//
	//}
	return res
}
