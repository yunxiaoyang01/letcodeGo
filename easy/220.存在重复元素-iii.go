/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

    for i:=0;i<len(nums);i++{
        for j:=i+1;j<=k-i;j++{
            if abs(nums[i],nums[j])<=t && abs(i,j)<=k{
                return true
            }
        }
    }
    return false
}

func abs(a,b int) int{
    if a-b<0{
        return -(a-b)
    }
    return a-b
}


// @lc code=end

