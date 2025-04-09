package main

import "strings"

func main() {
	isFliped := isFlipedString("waterbottle", "erbottlewat")
	println(isFliped)
}

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
