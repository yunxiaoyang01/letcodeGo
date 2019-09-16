package main

//letCode 盛最多水的容器
func main() {

}
func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			y := minInt(height[i], height[j])
			x := j - i
			if maxArea < x*y {
				maxArea = x * y
			}
		}
	}
	return maxArea
}
func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
