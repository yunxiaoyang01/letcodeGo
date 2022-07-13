package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%v", isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a, b := 0, x

	for b > 0 {
		a = a*10 + b%10
		b /= 10
	}
	if x == a {
		return true
	}
	return false
}
