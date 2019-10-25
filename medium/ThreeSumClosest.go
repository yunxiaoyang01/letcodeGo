package medium

import "math"

func main() {
	var nums = []int{-1, 2, 1, -4}
	target := 1
	println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int { //暴力解法
	n := len(nums)
	if n < 3 {
		return 0
	}
	diff := math.MaxInt32
	minTarge := 0
	if n == 3 {
		return minTarge
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abc(nums[i]+nums[j]+nums[k]-target) < diff {
					diff = abc(nums[i] + nums[j] + nums[k] - target)
					minTarge = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return minTarge
}

func abc(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
