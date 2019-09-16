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
