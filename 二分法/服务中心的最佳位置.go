package main

import (
	"fmt"
	"sort"
)

var list [][]int

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var ints []int
		var a, b int
		fmt.Scan(&a, &b)
		ints = append(ints, a, b)
		list = append(list, ints)
	}

	// 对线段进行排序
	// 左边界小的在前
	// 左边界相等的右边界小的在前
	sort.Slice(list, func(i, j int) bool {
		if list[i][0] == list[j][0] {
			return list[i][1] < list[j][1]
		}
		return list[i][0] < list[j][0]
	})

	// 最左边的坐标
	min := list[0][0]
	// 最右边的坐标
	max := list[n-1][1]
	// 距离最小值
	res := 1<<31 - 1
	for i := min; i <= max; i++ {
		res = Min(res, handle(i))
	}

	fmt.Println(res)
}

// 求出服务中心到所有街道的距离和
func handle(mid int) int {
	res := 0
	for _, ints := range list {
		a := ints[0]
		b := ints[1]
		if mid < a {
			res += a - mid
		} else if mid > b {
			res += mid - b
		}
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
