package main

func main() {

}

func totalHammingDistance(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += (val >> i) & 1
		}
		ans += c * (n - c)
	}
	return ans
}
