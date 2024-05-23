package main

func main() {
	nums := []int{1, 3, 2, 3, 1, 3}
	k := 3
	println(longestEqualSubarray(nums, k))
}

func longestEqualSubarray(nums []int, k int) int {
	pos := make(map[int][]int)
	for i, num := range nums {
		pos[num] = append(pos[num], i)
	}
	ans := 0
	for _, vec := range pos {
		j := 0
		for i := 0; i < len(vec); i++ {
			/* 缩小窗口，直到不同元素数量小于等于 k */
			for vec[i]-vec[j]-(i-j) > k {
				j++
			}
			ans = max(ans, i-j+1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
