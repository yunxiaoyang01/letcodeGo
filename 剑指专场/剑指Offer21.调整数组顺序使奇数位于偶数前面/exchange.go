package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("%+v", exchange(nums))
}

func exchange(nums []int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast]%2 != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	return nums
}
