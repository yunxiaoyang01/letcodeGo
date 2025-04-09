package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	target := 6
	fmt.Printf("%v\n", twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
			continue
		} else if sum < target {
			l++
			continue
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}
