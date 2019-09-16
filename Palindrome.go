package main

import "strconv"

func main() {

}
func isPalindrome(x int) bool {
	xstr := strconv.Itoa(x)
	xDao := reverseString(xstr)
	if xstr == xDao {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
