package easy

import "sort"

func main() {
	nums := []int{3, 2, 3}
	println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	max := 1
	sum := 1
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			if max < sum {
				max = sum
				result = nums[i]
			}
			sum = 1
		} else {
			sum++
			if max < sum {
				max = sum
				result = nums[i]
			}
		}
	}
	return result
}
func majorityElementEasy(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
