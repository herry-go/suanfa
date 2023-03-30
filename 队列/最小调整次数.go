package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	deque := list.New()
	index := 1
	res := 0
	for i := 0; i < 2*n; i++ {
		scanner.Scan()
		input := scanner.Text()
		if input == "remove" {
			if deque.Front().Value.(int) != index {
				// Convert deque to slice and sort in ascending order
				cp := make([]int, deque.Len())
				j := 0
				for e := deque.Front(); e != nil; e = e.Next() {
					cp[j] = e.Value.(int)
					j++
				}
				sort.Ints(cp)
				// Clear deque and add sorted slice back to deque
				deque.Init()
				for _, v := range cp {
					deque.PushBack(v)
				}
				res++
			}
			deque.Remove(deque.Front())
			index++
		} else {
			num, _ := strconv.Atoi(strings.Split(input, " ")[2])
			if strings.Split(input, " ")[0] == "head" {
				deque.PushFront(num)
			} else {
				deque.PushBack(num)
			}
		}
	}

	fmt.Println(res)
}