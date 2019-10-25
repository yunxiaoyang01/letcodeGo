package easy

func main() {
	/*nums :=[]int{2,2,1}
	println(singleNumber1(nums))*/
	println(5 ^ 5)
}
func singleNumber(nums []int) int {
	result := 0
	for _, x := range nums {
		result = result ^ x

	}
	return result
}
func singleNumber1(nums []int) int {
	re := map[int]int{}
	for _, x := range nums {
		re[x]++
	}
	result := 0
	for k, v := range re {
		if v == 1 {
			result = k
		}
	}
	return result
}
