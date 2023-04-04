package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	arr := strings.Split(line, " ")
	max := 0
	m := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		if i >= count {
			key := arr[i-count]
			m[key]--
			if m[key] == 0 {
				delete(m, key)
			}
		}
		m[arr[i]]++
		if m[arr[i]] > max {
			max = m[arr[i]]
		}
	}
	fmt.Println(max)
}
