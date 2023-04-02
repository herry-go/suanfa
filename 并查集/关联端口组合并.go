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
	var M int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &M)

	if M > 10 {
		fmt.Println("[[]]")
	} else {
		parent := make([]int, M)
		for i := range parent {
			parent[i] = i
		}

		for i := 0; i < M; i++ {
			scanner.Scan()
			input := scanner.Text()
			nums := strings.Split(input, ",")
			if len(nums) == 1 {
				continue
			}
			for j := 1; j < len(nums); j++ {
				num1, _ := strconv.Atoi(nums[j])
				num2, _ := strconv.Atoi(nums[j-1])
				union(parent, num1-1, num2-1)
			}
		}

		sets := make(map[int]map[int]bool)
		for i := 0; i < M; i++ {
			root := find(parent, i)
			if _, ok := sets[root]; !ok {
				sets[root] = make(map[int]bool)
			}
			scanner.Scan()
			nums := strings.Split(scanner.Text(), ",")
			for _, numStr := range nums {
				num, _ := strconv.Atoi(numStr)
				sets[root][num] = true
			}
		}

		var result []map[int]bool
		for _, set := range sets {
			if len(set) >= 2 {
				result = append(result, set)
			}
		}

		fmt.Println(result)
	}
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent []int, i int, j int) {
	parent[find(parent, i)] = find(parent, j)
}
