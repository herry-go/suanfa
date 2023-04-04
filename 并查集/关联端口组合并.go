package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var m int
	fmt.Scan(&m)

	parents := make(map[int]int)
	for i := 0; i < m; i++ {
		var input string
		fmt.Scan(&input)

		var parent int
		for _, portStr := range strings.Split(input, ",") {
			port,_ := strconv.Atoi(portStr)

			if _, ok := parents[port]; !ok {
				parents[port] = port
			}
			if parent == 0 {
				parent = parents[port]
			} else {
				parents[find(parents, port)] = parent
			}
		}
	}

	mergedGroups := make(map[int][]int)
	for port, parent := range parents {
		mergedGroups[find(parents, parent)] = append(mergedGroups[find(parents, parent)], port)
	}

	result := make([][]int, 0)
	for _, group := range mergedGroups {
		sort.Ints(group)
		result = append(result, group)
	}

	fmt.Println(result)
}

func find(parents map[int]int, port int) int {
	if parents[port] == port {
		return port
	}
	parents[port] = find(parents, parents[port])
	return parents[port]
}
