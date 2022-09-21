package main

import (
	"fmt"
	"sort"
	"strings"
)

//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.

func searchAnagram(sl []string) map[string][]string {

	m1 := make(map[string][]string)

	for _, str := range sl {
		runes := []rune(strings.ToLower(str))
		//var newStr = ""

		sort.Slice(runes, func(i, j int) bool {
			if string(runes[i]) < string(runes[j]) {
				return true
			}
			return false
		})
		_, ok := m1[string(runes)]

		if !ok {
			arr := make([]string, 1)
			arr[0] = strings.ToLower(str)
			m1[string(runes)] = arr
		} else {
			m1[string(runes)] = append(m1[string(runes)], strings.ToLower(str))
		}
		//fmt.Println("rune = ", string(runes), i)
	}
	//fmt.Println("map = ", m1)

	resMap := make(map[string][]string)

	for _, slice := range m1 {
		if len(slice) > 1 {
			resMap[slice[0]] = slice
		}
	}

	fmt.Println("resMap = ", resMap)

	return resMap
}

func main() {
	sl := []string{"абфыяддыАбаЛБМТВФвйуц", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	searchAnagram(sl)

}
