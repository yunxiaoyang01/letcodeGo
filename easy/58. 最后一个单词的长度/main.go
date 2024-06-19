package main

import "strings"

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != " " && arr[i] != "" {
			return len(arr[i])
		}
	}
	return 0
}
