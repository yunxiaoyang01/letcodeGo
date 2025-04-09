package main

import "fmt"

func main() {
	digits := []int{9, 9, 9, 9, 9}
	fmt.Printf("%v", plusOne(digits))
}

func plusOne(digits []int) []int {
	add := false
	target := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]
		if i == len(digits)-1 {
			num += 1
		}
		if add {
			num += 1
		}
		if num >= 10 {
			add = true
			num = num % 10
		} else {
			add = false
		}
		target = append(target, num)
	}
	if add {
		target = append(target, 1)
	}
	return reverse(target)
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
