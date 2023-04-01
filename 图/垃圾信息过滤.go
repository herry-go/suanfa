package main

import "fmt"

func main() {
	var n, from, to, id, l, m1, m2 int
	var res bool
	data := make([][]int, 100)
	for i := range data {
		data[i] = make([]int, 100)
	}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&from, &to)
		data[from][to]++
	}
	fmt.Scan(&id)
	for i := 0; i < 100; i++ {
		if data[id][i] > 0 && data[i][id] == 0 {
			l++
		}
		m1 += data[id][i]
		m2 += data[i][id]
		if data[id][i]-data[i][id] > 5 {
			res = true
		}
	}
	if l > 5 || m1-m2 > 10 {
		res = true
	}
	fmt.Printf("%v %v %v", res, l, m1-m2)
}
