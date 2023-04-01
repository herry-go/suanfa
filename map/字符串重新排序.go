package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type StrSort struct {
	value string
	count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		var maps = make(map[string]int)
		strs := strings.Split(scanner.Text(), " ")
		for _, v := range strs {
			s := strings.Split(v, "")
			sort.Strings(s)
			str := strings.Join(s, "")
			maps[str]++
		}
		var list = make([]StrSort, 0)
		for k, v := range maps {
			s := StrSort{
				value: k,
				count: v,
			}
			list = append(list, s)
		}

		sort.Slice(list, func(i, j int) bool {
			if list[i].count > list[j].count {
				return true
			} else if list[i].count == list[j].count {
				if len(list[i].value) < len(list[j].value) {
					return true
				} else if len(list[i].value) == len(list[j].value) {
					return list[i].value < list[j].value
				}
			}
			return false
		})

		var value = ""
		for _, v := range list {
			for i := 0; i < v.count; i++ {
				value += v.value + " "
			}
		}
		fmt.Println(value)

	}

}
