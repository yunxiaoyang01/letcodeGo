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
	i, j := 0, len(s)-1
	s = strings.ToLower(s)
	for i < j {
		if !(('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9')) {
			i++
			continue
		}
		if !(('a' <= s[j] && s[j] <= 'z') || ('0' <= s[j] && s[j] <= '9')) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
