package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := []int{}
	for x > 0 {
		t := x % 10
		ans = append(ans, t)
		x = x / 10
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		if ans[i] != ans[j] {
			return false
		}
	}
	return true
}
