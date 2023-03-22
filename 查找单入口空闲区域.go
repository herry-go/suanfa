package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hanShu int
var lieShu int
var count int
var rukou = make([]int, 2)
var list [][]string

func main() {

	fmt.Scan(&hanShu, &lieShu)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < hanShu; i++ {
		var han string
		scanner.Scan()
		han = scanner.Text()
		hanList := strings.Split(han, " ")
		list = append(list, hanList)
	}
	fmt.Println(list)

	var max int = 0
	var quyu = [][]int{}
	for hs, h := range list {
		fmt.Println("行数：", hs)
		for ls, v := range h {
			fmt.Println("列数：", hs, v)
			if list[hs][ls] == "o" {
				list[hs][ls] = "x"
				zuobiao := [][]int{}
				zuobiao = append(zuobiao, []int{hs, ls})
				rukouShu(ls, ls, &zuobiao)
				if count == 1 {
					if max == len(zuobiao) {
						quyu = [][]int{}
					} else if max < len(zuobiao) {
						quyu = [][]int{}
						quyu = append(quyu, []int{rukou[0], rukou[1], len(zuobiao)})
						max = len(zuobiao)
					}
				}
				count = 0
				rukou = make([]int, 2)
			}
		}
	}

	if len(quyu) == 1 {
		res := quyu[0]
		fmt.Println(fmt.Sprintf("%d %d %d", res[0], res[1], res[2]))
	} else if max != 0 {
		fmt.Println(max)
	} else {
		fmt.Println("NULL")
	}

}

func rukouShu(hs int, ls int, zuobiao *[][]int) {
	// 表示坐标在边界
	if hs == 0 || hs == hanShu-1 || ls == 0 || ls == lieShu-1 {
		count++
		rukou[0] = hs
		rukou[1] = ls
	}

	if hs < hanShu-1 {
		if list[hs+1][ls] == "o" {
			list[hs+1][ls] = "x"
			*zuobiao = append(*zuobiao, []int{hs + 1, ls})
			rukouShu(hs+1, ls, zuobiao)
		}
	}

	if ls < lieShu-1 {
		if list[hs][ls+1] == "o" {
			list[hs][ls+1] = "x"
			*zuobiao = append(*zuobiao, []int{hs, ls + 1})
			rukouShu(hs, ls+1, zuobiao)
		}
	}

	if hs < hanShu-1 && ls > 0 {
		if list[hs][ls-1] == "o" {
			list[hs][ls-1] = "x"
			*zuobiao = append(*zuobiao, []int{hs, ls - 1})
			rukouShu(hs, ls-1, zuobiao)
		}
	}

}
