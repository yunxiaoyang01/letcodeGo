package easy

func main() {
	println(climbStairs(45))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	array := []int{1, 2}
	if n > 2 {
		for i := 2; i < n; i++ {
			result := array[i-1] + array[i-2]
			array = append(array, result)
		}
	}
	return int(array[n-1])
}
