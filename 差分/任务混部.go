package main

import (
	"fmt"
)

func main() {
	var taskNum int
	fmt.Scan(&taskNum)

	serversMap := make([]int, 50000)

	for i := 0; i < taskNum; i++ {
		var startTime, endTime, parallelism int
		fmt.Scan(&startTime, &endTime, &parallelism)
		serversMap[startTime] += parallelism
		serversMap[endTime] -= parallelism
	}

	totalServerNum := 0
	res := 0
	for i := 0; i < len(serversMap); i++ {
		totalServerNum += serversMap[i]
		if totalServerNum > res {
			res = totalServerNum
		}
	}

	fmt.Println(res)
}

//0 2  3  5 6 9
//1 1 -1 -1 2 -2
