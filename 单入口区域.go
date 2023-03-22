package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n int
		fmt.Sscanf(scanner.Text(), "%d %d", &m, &n)
		ca := make([][]string, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			ca[i] = strings.Split(scanner.Text(), " ")
		}

		maxCount := 0
		maps := make(map[string]int)
		for j := 0; j < n; j++ {
			if ca[0][j] == "o" {
				count := calc(copy(ca), 0, j, true)
				if count > 0 {
					key := fmt.Sprintf("%d %d", 0, j)
					maps[key] = count
					if count > maxCount {
						maxCount = count
					}
				}
			}

			if ca[m-1][j] == "o" {
				count2 := calc(copy(ca), m-1, j, true)
				if count2 > 0 {
					key := fmt.Sprintf("%d %d", m-1, j)
					maps[key] = count2
					if count2 > maxCount {
						maxCount = count2
					}
				}
			}
		}

		for i := 1; i < m-1; i++ {
			if ca[i][0] == "o" {
				count := calc(copy(ca), i, 0, true)
				if count > 0 {
					key := fmt.Sprintf("%d %d", i, 0)
					maps[key] = count
					if count > maxCount {
						maxCount = count
					}
				}
			}

			if ca[i][n-1] == "o" {
				count2 := calc(copy(ca), i, n-1, true)
				if count2 > 0 {
					key := fmt.Sprintf("%d %d", i, n-1)
					maps[key] = count2
					if count2 > maxCount {
						maxCount = count2
					}
					fmt.Println(maxCount)
				}

			}

			var maxKey string = ""
			for key, value := range maps {
				if value == maxCount {
					if maxKey == "" {
						maxKey = key
					} else {
						maxKey = "more"
						break
					}
				}
			}

			if maxCount == 0 {
				fmt.Println("NULL")
			} else if maxKey == "more" {
				fmt.Println(maxCount)
			} else {
				fmt.Println(maxKey, maxCount)
			}

		}

	}

}

func calc(ca [][]string, i int, j int, isRuKou bool) int {
	if !isRuKou {
		if i == 0 || i == len(ca)-1 || j == 0 || j == len(ca[0])-1 {
			return 0
		}
	}

	ca[i][j] = "x"
	count := 1

	if j+1 < len(ca[0]) && ca[i][j+1] == "o" {
		count1 := calc(ca, i, j+1, false)
		if count1 == 0 {
			return 0
		}
		count += count1
	}

	if i+1 < len(ca) && ca[i+1][j] == "o" {
		count1 := calc(ca, i+1, j, false)
		if count1 == 0 {
			return 0
		}
		count += count1
	}

	if j-1 >= 0 && ca[i][j-1] == "o" {
		count1 := calc(ca, i, j-1, false)
		if count1 == 0 {
			return 0
		}
		count += count1
	}

	if i-1 >= 0 && ca[i-1][j] == "o" {
		count1 := calc(ca, i-1, j, false)
		if count1 == 0 {
			return 0
		}
		count += count1
	}

	return count
}

func copy(ca [][]string) [][]string {
	m := len(ca)
	n := len(ca[0])
	newCa := make([][]string, m)
	for i := 0; i < m; i++ {
		newCa[i] = make([]string, n)
		for j := 0; j < n; j++ {
			newCa[i][j] = ca[i][j]
		}
	}
	return newCa
}
