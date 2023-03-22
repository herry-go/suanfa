package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var all float64

func main() {
	var han int
	fmt.Scan(&han)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < han; i++ {
		scanner.Scan()
		input := scanner.Text()
		t := ""
		c := ""
		for _, x := range []rune(input) {
			if unicode.IsDigit(x) {
				if t != "" {
					s, _ := strconv.Atoi(c)
					all += huansuan(t, s)
					t = ""
					c = ""
				}
				c += string(x)
			} else {
				t += string(x)
			}
		}
		s, _ := strconv.Atoi(c)
		all += huansuan(t, s)

	}
	fmt.Println(int(all))
}

func huansuan(t string, num int) (res float64) {

	switch t {
	case "CNY":
		res = float64(num) * 100
	case "fen":
		res = float64(num)
	case "JPY":
		res = float64(num) / 1825 * 10000
	case "sen":
		res = float64(num) / 1825 * 100
	case "HKD":
		res = float64(num) / 123 * 10000
	case "cents":
		res = float64(num) / 123 * 100
	case "EUR":
		res = float64(num) / 14 * 10000
	case "eurocents":
		res = float64(num) / 14 * 100
	case "GBP":
		res = float64(num) / 12 * 10000
	case "pence":
		res = float64(num) / 12 * 100
	}
	return res
}
