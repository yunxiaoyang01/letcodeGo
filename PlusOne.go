package main

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	result := plusOne(digits)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
func plusOne(digits []int) []int {
	var result []int
	var fanResult []int
	flag := false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] = digits[i] + 1
		}
		if digits[i] > 9 {
			digits[i] = digits[i] - 10
			if i != 0 {
				digits[i-1] = digits[i-1] + 1
			} else {
				flag = true
			}
		}
		result = append(result, digits[i])
	}
	if flag {
		result = append(result, 1)
	}
	for i := len(result) - 1; i >= 0; i-- {
		fanResult = append(fanResult, result[i])
	}
	return fanResult
}

func plusOneTop(digits []int) []int { //精妙算法
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
