package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m int
var n int
var Nums [][]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &m, &n)

		Nums = make([][]int, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			strs := strings.Split(scanner.Text(), " ")
			list := make([]int, n)
			for j, v := range strs {
				d, _ := strconv.Atoi(v)
				list[j] = d
			}
			Nums[i] = list
		}

		var maps = make(map[string][]int)
		for i, list := range Nums {
			for j, _ := range list {
				var dst = make([][]int, len(Nums))
				copy(dst, Nums)
				quyu := findquyu(i, j, dst)
				maps[fmt.Sprintf("%d-%d", i, j)] = quyu
			}
		}
		var max int
		var maxk string
		for k, v := range maps {
			if len(v) > max {
				max = len(v)
				maxk = k
			}
		}
		fmt.Println(max + 1)
		fmt.Println(maxk)

	}
}

func findquyu(i int, j int, nums [][]int) (res []int) {

	if i == m-1 && j == n-1 {
		return res
	}
	num := nums[i][j]

	if num == -1 {
		return nil
	}
	nums[i][j] = -1

	//上
	if i > 0 && jdz(num, nums[i-1][j]) {
		res = append(res, nums[i-1][j])
		r := findquyu(i-1, j, nums)
		if len(r) > 0 {
			res = append(res, r...)
		}
	}
	//下
	if i < len(nums)-1 && jdz(num, nums[i+1][j]) {
		res = append(res, nums[i+1][j])
		r := findquyu(i+1, j, nums)
		if len(r) > 0 {
			res = append(res, r...)
		}
	}
	//左
	if j > 0 && jdz(num, nums[i][j-1]) {
		res = append(res, nums[i][j-1])
		r := findquyu(i, j-1, nums)
		if len(r) > 0 {
			res = append(res, r...)
		}
	}
	//右
	if j < len(nums[0])-1 && jdz(num, nums[i][j+1]) {
		res = append(res, nums[i][j+1])
		r := findquyu(i, j+1, nums)
		if len(r) > 0 {
			res = append(res, r...)
		}
	}

	return res
}

func jdz(a, b int) bool {
	if a > b {
		return a-b <= 1
	} else if a < b {
		return b-a <= 1
	} else {
		return true
	}
}
