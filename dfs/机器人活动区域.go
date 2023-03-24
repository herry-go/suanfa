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
	for scanner.Scan() {
		var m int
		var n int
		nums := make([][]int, m)
		fmt.Sscanf(scanner.Text(), "%d %d", &m, &n)
		for i := 0; i < m; i++ {
			scanner.Scan()
			strs := strings.Split(scanner.Text(), " ")
			list := make([]int, n)
			for j, v := range strs {
				d, _ := strconv.Atoi(v)
				list[j] = d
			}
			nums[i] = list
		}
		fmt.Println(nums)
	}
}
