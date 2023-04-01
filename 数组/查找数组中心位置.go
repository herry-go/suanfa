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
		strs := strings.Split(scanner.Text(), " ")

		var nums []int
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}

		var left = 1
		var right = 1
		var flag = false
		for _, v := range nums {
			right *= v
		}

		for i, v := range nums {
			if i > 0 {
				left *= nums[i-1]
			}
			right = right / v
			if left == right {
				fmt.Println(i)
				flag = true
				break
			}

		}

		if !flag {
			fmt.Println(-1)
		}
	}

}
