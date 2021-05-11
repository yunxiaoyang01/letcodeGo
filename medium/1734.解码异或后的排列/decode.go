package main

import "fmt"

func main() {
	encode := []int{3, 1}
	fmt.Printf("%d", decode(encode))
}

func decode(encoded []int) []int {
	n := len(encoded)
	x := 0
	for i := 1; i <= n+1; i++ {
		x ^= i
	}
	y := 0
	for i := 1; i < n; i += 2 {
		y ^= encoded[i]
	}
	perm := make([]int, n+1)
	perm[0] = x ^ y
	for i := 0; i < n; i++ {
		perm[i+1] = perm[i] ^ encoded[i]
	}
	return perm
}
