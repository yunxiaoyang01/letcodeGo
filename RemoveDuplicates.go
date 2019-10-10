package main

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(nums, 2))
}

func removeDuplicates(nums []int) int {
	k := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[k] {
			nums[k+1] = nums[j]
			k++
		}
	}
	return k + 1
}
func removeElement(nums []int, val int) int {
	k := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[k] = nums[j]
			k++
		}
	}
	return k
}
