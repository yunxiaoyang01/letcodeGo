package main

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for _, x := range rotate(nums, 3) {
		println(x)
	}

}

func rotate(nums []int, k int) []int {
	k %= len(nums)
	ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	nums = append(nums[:0], ans...)
	return nums
}
