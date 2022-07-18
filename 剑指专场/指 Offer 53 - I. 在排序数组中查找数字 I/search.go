package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	s := map[int]int{}
	for _, num := range nums {
		s[num]++
	}
	return s[target]
}
