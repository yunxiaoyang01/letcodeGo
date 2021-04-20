/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow :=0
	fast :=1
	for ;fast<len(nums);{
		if nums[slow]!=nums[fast]{
			slow +=1
			nums[slow]=nums[fast]
		}
		fast+=1
	}
	return slow+1
}
// @lc code=end

