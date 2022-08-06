package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	newStr := ""
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			newStr += string(v)
		}
	}
	lowStr := strings.ToLower(newStr)
	left, right := 0, len(lowStr)-1
	for left < right {
		if lowStr[left] == lowStr[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
