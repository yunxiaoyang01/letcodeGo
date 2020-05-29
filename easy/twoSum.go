package easy

//func main() {
//	numbers := []int{2, 7, 11, 15}
//	target := 9
//	for _, x := range twoSum(numbers, target) {
//		println(x)
//	}
//}

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			break
		}
	}
	result := []int{l + 1, r + 1}
	return result
}
