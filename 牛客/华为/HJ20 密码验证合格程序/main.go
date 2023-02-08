package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		if valid(str) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("NG\n")
		}
	}

}

func valid(str string) bool {
	// rule1
	if len(str) < 8 {
		return false
	}
	// rule2
	dictMap := map[int]int{}
	for i := 0; i < len(str); i++ {
		if '0' <= str[i] && str[i] <= '9' {
			dictMap[0]++
		} else if 'a' <= str[i] && str[i] <= 'z' {
			dictMap[1]++
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			dictMap[2]++
		} else {
			dictMap[3]++
		}
	}
	count := 0
	for _, v := range dictMap {
		if v != 0 {
			count++
		}
	}
	if count < 3 {
		return false
	}
	// rule3
	repeatMap := map[string]struct{}{}
	for i := 0; i < len(str)-2; i++ {
		if _, ok := repeatMap[str[i:i+3]]; ok {
			return false
		}
		repeatMap[str[i:i+3]] = struct{}{}
	}
	return true
}
