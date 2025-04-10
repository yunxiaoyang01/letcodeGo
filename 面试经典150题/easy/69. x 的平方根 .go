package main

import "fmt"

func main() {

	fmt.Println(mySqrtV2(256))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := 1
	for i := 1; i <= x/2; i++ {
		if i*i <= x {
			ans = i
		} else {
			break
		}
	}
	return ans
}

func mySqrtV2(x int) int {
	l, r := 0, x/2
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
