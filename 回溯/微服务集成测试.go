package main

import (
	"fmt"
)

func main() {

	for {
		var n int
		fmt.Scan(&n)
		nums := make([][]int, n)
		for i := 0; i < n; i++ {
			nums[i] = make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scan(&nums[i][j])
			}
		}
		var k int
		fmt.Scan(&k)
		count := 0
		count = dfs(nums, k-1)
		fmt.Println(count)
	}

}

func dfs(nums [][]int, k int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[k][i] != 0 && i != k {
			temp := dfs(nums, i)
			if temp > max {
				max = temp
			}
		}
	}
	fmt.Println(max)
	return max + nums[k][k]
}
