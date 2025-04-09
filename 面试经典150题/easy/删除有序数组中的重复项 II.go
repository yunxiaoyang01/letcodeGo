package main

func main() {

}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	left, right := 2, 2
	for right < n {
		if nums[left-2] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
