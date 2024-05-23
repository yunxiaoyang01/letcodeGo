package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(movesToMakeZigzag(nums))
}

/*
提示 1
只有减少操作意味着什么？
如果你想把 x 小于等于 y 换成 x > y 只能减少y，x是没有用的
提示 2
题目给了两种锯齿数组，分别考虑
提示 3-1
以第一种为例，要想转换成 A[0] > A[1] < A[2] > A[3] < A[4] 根据提示 1，哪些数一定不需要减少？
提示 3-2
nums[0],nums[2],⋯ 是一定不需要的。
你也可以这样思考：
假设把 nums 转换成第一种锯齿形，且
nums[2] 减少了，那么把 nums[2] 恢复成它的原始值，数组仍然符合第一种锯齿形。所以 nums[2] 是不需要减少的。其它数同理。
提示 4
为了使操作次数尽量少，nums[i] 不断减小到要比左右相邻数字都小，就立刻停止，所以 nums[i] 要修改成
m=min(nums[i−1],nums[i+1])−1 修改次数为 nums[i]−m
如果 nums[i] 本来就不超过 m，就无需修改。
因此，nums[i] 的修改次数为
max(nums[i]−min(nums[i−1],nums[i+1])+1,0)
如果 i−1 或者 i+1 下标越界，则对应的数字视作无穷大。
最后，把奇数偶数下标对应的修改次数分别累加，结果分别设为 s1和s2，最终求 min(s1,s2)

*/

func movesToMakeZigzag(nums []int) int {
	s := [2]int{}
	for i, x := range nums {
		left, right := math.MaxInt, math.MaxInt
		if i > 0 {
			left = nums[i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1]
		}
		s[i%2] += max(x-min(left, right)+1, 0)
	}
	return min(s[0], s[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
