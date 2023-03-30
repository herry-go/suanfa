package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var str string
		var wordStr string

		str = scanner.Text()
		scanner.Scan()
		wordStr = scanner.Text()

		words := toArrayStr(wordStr)
		var left int = 0
		var right int = len(words)
		var count int = 0
		for {
			if right > len([]rune(str)) {
				break
			}

			sub := str[left:right]
			subs := toArrayStr(sub)
			var flag bool = true
			for i, v := range subs {
				if words[i] != v {
					flag = false
				}
			}

			if flag {
				count++
			}

			left++
			right++
		}
		fmt.Println(count)

	}

}

func toArrayStr(str string) []string {
	var subs []string
	for _, v := range []rune(str) {
		subs = append(subs, string(v))
	}
	sort.Strings(subs)
	return subs
}
