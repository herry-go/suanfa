package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		words := strings.Split(scanner.Text(), " ")
		wordCount := make(map[string]int)
		for _, word := range words {
			chars := strings.Split(word, "")
			sort.Strings(chars)
			sortedWord := strings.Join(chars, "")
			wordCount[sortedWord]++
		}
		type kv struct {
			Key   string
			Value int
		}
		var sortedWords []kv
		for k, v := range wordCount {
			sortedWords = append(sortedWords, kv{k, v})
		}
		sort.Slice(sortedWords, func(i, j int) bool {
			if sortedWords[i].Value > sortedWords[j].Value {
				return true
			} else if sortedWords[i].Value == sortedWords[j].Value {
				if len(sortedWords[i].Key) < len(sortedWords[j].Key) {
					return true
				} else if len(sortedWords[i].Key) == len(sortedWords[j].Key) {
					return sortedWords[i].Key < sortedWords[j].Key
				}
			}
			return false
		})
		var res string
		for _, word := range sortedWords {
			for i := 0; i < word.Value; i++ {
				res += word.Key + " "
			}
		}
		fmt.Println(strings.TrimSuffix(res, " "))

	}
}
