package anagramm

import (
	"sort"
	"strings"
)

/*
	Алгоритм
	Иду по слайсу
	В мапу как ключ кладу отсортированную и toLower строку.
	Далее я сравниваю каждое слово с ключём(то есть если есть m[отсортированное слово] есть, то я кладу туда это слово без сортировки)

	Вопрос:
	Как вернуть слово-ключ в нормальное состояние
	- Хранить в массиве или создать еще одну мапу в которую все перекинуть все значени, но уже по нормальным ключам
*/

func FindAnagram(str *[]string) map[string][]string {
	tmpMap := make(map[string][]string)
	nameArray := make([]string, 0)
	nameArraySorted := make([]string, 0)
	sliceString := make([]string, 0)
	tmpString := ""
	for _, str := range *str {
		str = strings.ToLower(str)
		sliceString = strings.Split(str, "")
		sort.Strings(sliceString)
		tmpString = concatenate(sliceString)
		_, ok := tmpMap[tmpString]
		if ok {
			tmpMap[tmpString] = append(tmpMap[tmpString], str)
		} else {
			nameArray = append(nameArray, str)
			nameArraySorted = append(nameArraySorted, tmpString)
			a := make([]string, 0)
			tmpMap[tmpString] = a
		}
	}
	refactorMap(&tmpMap, nameArray, nameArraySorted)
	return tmpMap
}

/*
	Сравниваю два строки на наличие всех символов и на одинаковую длину
*/
func refactorMap(ma *map[string][]string, strOrder []string, strChaos []string) {
	for i, s := range strOrder {
		if _, ok := (*ma)[s]; !ok {
			(*ma)[s] = (*ma)[strChaos[i]]
			delete(*ma, strChaos[i])
		}
	}
}

func concatenate(str []string) string {
	res := ""
	for _, s := range str {
		res += s
	}
	return res
}
