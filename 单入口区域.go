package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 这是一个Go程序，从标准输入读取输入并进行处理。程序从标准输入读取字符矩阵，并查找与矩阵边界相连的最大连通分量的“o”字符。然后程序输出此连通分量的大小
// 和其中一个单元格的坐标。
//
// 程序使用扫描仪从标准输入读取输入，并在主函数中进行处理。主函数从标准输入读取矩阵并将其存储在2D字符串切片中。然后函数遍历矩阵边界上的单元格，并在包
// 含“o”的每个单元格上调用calc函数。calc函数递归地探索包含输入单元格的“o”单元格的连通分量，并返回此组件的大小。主函数将最大连通分量的大小存储在
// maxCount变量中，并将所有大小与最大组件相同的连通分量的大小存储在映射中。然后主函数遍历映射中的键，并输出最大连通分量的大小和其中一个单元格的坐标。如果有多个大小相同的连通分量，则程序仅输出最大组件的大小。
//
// calc函数以矩阵的副本、单元格的坐标和一个布尔标志作为输入，该标志指示单元格是否在矩阵的边界上。该函数递归地探索包含输入单元格的“o”单元格的连通分量，并返回此组件的大小。该函数使用“x”标记每个探索的单元格，以避免再次探索它。
//
// copy函数以矩阵作为输入并返回矩阵的副本。该函数用于在调用calc函数之前创建矩阵的副本，以避免修改原始矩阵。
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
