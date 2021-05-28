package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	ans := 0
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return ans
}
