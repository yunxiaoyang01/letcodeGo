package main

func main() {

}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(area, ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
