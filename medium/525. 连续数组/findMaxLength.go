package main

import "fmt"

func main() {
	nums := []int{0, 1, 0}
	fmt.Printf("%d", findMaxLength(nums))
}

func findMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	count := 0
	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if prefix, has := mp[count]; has {
			maxLength = max(maxLength, i-prefix)
		} else {
			mp[count] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
