package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	var nums = []int{}
	for _, v := range strings.Split(input, ","){
		port,_ := strconv.Atoi(v)
		nums = append(nums, port)
	}

	m := longestConsecutive(nums)
	fmt.Println(m)
}

func longestConsecutive(nums []int) int {
	numMap := make(map[int]int)
	max_length := 0

	for _, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = num

			if _, ok := numMap[num-1]; ok {
				union(numMap, num, num-1)
			}

			if _, ok := numMap[num+1]; ok {
				union(numMap, num, num+1)
			}
		}
	}

	for _, parent := range numMap {
		if parent == numMap[parent] {
			curr_num := parent
			curr_length := 1

			for numMap[curr_num+1] != curr_num+1 {
				curr_num += 1
				curr_length += 1
			}

			max_length = max(max_length, curr_length)
		}
	}

	return max_length
}

func find2(numMap map[int]int, num int) int {
	if numMap[num] != num {
		numMap[num] = find2(numMap, numMap[num])
	}
	return numMap[num]
}

func union(numMap map[int]int, num1 int, num2 int) {
	root1 := find2(numMap, num1)
	root2 := find2(numMap, num2)
	numMap[root1] = root2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


