package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	i, j := 0, 0
	for i < sLen && j < tLen {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == sLen
}
