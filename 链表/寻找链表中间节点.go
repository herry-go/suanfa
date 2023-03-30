package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Node struct {
	data string
	next string
}

func main() {
	// 通过map快速寻址
	m := make(map[string]Node)
	scanner := bufio.NewScanner(os.Stdin)
	// 初始化元信息和具体数据
	scanner.Scan()
	mate := strings.Split(scanner.Text(), " ")
	for i := 0; i < parseInt(mate[1]); i++ {
		scanner.Scan()
		info := strings.Split(scanner.Text(), " ")
		m[info[0]] = Node{data: info[1], next: info[2]}
	}

	// 因为存在为空的数据，通过list保存有效数据和记录数据链长度
	var res []string
	address := mate[0]
	for {
		if _, ok := m[address]; !ok {
			break
		}
		// 记录当前节点，并指向下一个节点
		node := m[address]
		res = append(res, node.data)
		address = node.next
	}

	// 返回中间第二个（如果有）的值
	idx := len(res) / 2
	fmt.Println(res[idx])
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

