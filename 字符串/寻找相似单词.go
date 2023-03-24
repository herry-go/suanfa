package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m int
		fmt.Sscanf(scanner.Text(), "%d", &m)

		var strs []string
		for i := 0; i <= m; i++ {
			scanner.Scan()
			str := scanner.Text()
			list := []string{}
			for _, v := range []rune(str) {
				list = append(list, string(v))
				sort.Strings(list)
			}
			strs = append(strs, strings.Join(list, ""))
		}

		var key string
		key = strs[len(strs)-1]

		for i, v := range strs {
			if key == v && i != len(strs)-1 {
				fmt.Println(v)
				return
			}
		}
		fmt.Println("NULL")
	}

}
