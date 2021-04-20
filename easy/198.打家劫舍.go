/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	length := len(nums)

	if length==1{
		return nums[0]
	}
	if length==2{
		return max(nums[0],nums[1])
	}
	first,second := nums[0],max(nums[0],nums[1])
	for i:=2;i<length;i++{
		first,second = second,max(first+nums[i],second)
	}
	return second
}
func max(a,b int) int  {
	if a>b {
		return a
	}
	return b
}

// @lc code=end

