/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n<1{
		return false
	}
	if n==1{
		return true
	}
	if n%2==0{
		return isUgly(n/2)
	}
	if n%3==0{
		return isUgly(n/3)
	}
	if n%5==0{
		return isUgly(n/5)
	}
	return false
	
}
// @lc code=end

