package main

import "fmt"

func main() {
	s := "	"
	t := "ahbgdc"
	fmt.Printf("%v", isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	index := 0
	for i := range t {
		if t[i] == s[index] {
			index++
			if index == len(s) {
				return true
			}
		}
	}
	return false
}
