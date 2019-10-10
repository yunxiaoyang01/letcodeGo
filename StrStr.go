package main

import "strings"

func main() {
	haystack := "aaaaa"
	needle := "bba"
	println(strStr(haystack, needle))
}
func strStr(haystack string, needle string) int {
	len := strings.Index(haystack, needle)
	return len
}
