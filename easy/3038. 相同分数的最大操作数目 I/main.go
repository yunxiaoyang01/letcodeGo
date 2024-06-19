package main

func main() {
	nums := []int{3, 2, 1, 4, 5}
	print(maxOperations(nums))
}
func maxOperations(nums []int) int {
	n, t := len(nums), 0
	for i := 1; i < n; i += 2 {
		if nums[i]+nums[i-1] != nums[1]+nums[0] {
			break
		}
		t++
	}
	return t
}
