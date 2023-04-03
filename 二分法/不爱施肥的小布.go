package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fields := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&fields[i])
	}

	l := 1
	sort.Ints(fields)
	r := fields[m-1]
	for l < r {
		mid := (l + r) / 2
		if check(mid, fields) <= n {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if check(l, fields) > n {
		fmt.Println(-1)
	} else {
		fmt.Println(l)
	}
}

func check(mid int, fields []int) int {
	ans := 0
	for i := 0; i < len(fields); i++ {
		if fields[i]%mid == 0 {
			ans += fields[i] / mid
		} else {
			ans += fields[i]/mid + 1
		}
	}
	return ans
}
