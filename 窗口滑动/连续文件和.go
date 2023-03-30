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
		var sum int
		var nums []int = make([]int, 0)
		var line string

		fmt.Scan(&sum)
		scanner.Scan()
		line = scanner.Text()
		for _, v := range strings.Split(line, " ") {
			d, _ := strconv.Atoi(v)
			nums = append(nums, d)
		}

		var left = 0
		var right = 1
		var max = 0
		var count int = nums[0]
		for {

			if max == sum || (right >= len(nums) && count < max) {
				break
			}

			if count <= sum {
				if max < count {
					max = count
				}
				if right < len(nums) {
					count += nums[right]
					right++
				}

			} else {
				count -= nums[left]
				left++
			}

		}
		fmt.Println(count)

	}

}
