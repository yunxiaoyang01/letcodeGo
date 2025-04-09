package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	front, end := 0, len(nums)-1
	for front <= end {
		if nums[front] == val {
			nums[front] = nums[end]
			end--
		} else {
			front++
		}
	}
	return front
}
