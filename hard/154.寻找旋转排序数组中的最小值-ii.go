/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	l,r := 0,len(nums)-1
	for l<r{
		mid := (l+r)/2
		if (nums[mid]>nums[r]){
			l = mid+1
		}else if (nums[mid]<nums[r]){
			r =  mid
		}else {
			r = r-1
		}
	}
	return nums[l]
}
// @lc code=end

