package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	arr = filterSpace(arr)
	arr = reverse(arr)
	str := ""
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 {
			str += arr[i] + " "
		} else {
			str += arr[i]
		}
	}
	return str
}

func filterSpace(arr []string) []string {
	ans := []string{}
	for _, v := range arr {
		if v != "" {
			ans = append(ans, v)
		}
	}
	return ans
}

func reverse(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
