package main

import "fmt"

func main() {
	first := "ab"
	second := "bc"
	fmt.Printf("%v", oneEditAway(first, second))
}

func oneEditAway(first, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, v := range second {
		if first[i] != byte(v) {
			if m == n {
				return first[i+1:] == second[i+1:]
			} else {
				return first[i+1:] == second[i:]
			}
		}
	}
	return true
}
