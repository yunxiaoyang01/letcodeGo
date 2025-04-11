package main

import "fmt"

func main() {
	fmt.Println(intToRoman(58))
}

func intToRoman(num int) string {
	intArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(intArr); i++ {
		for num >= intArr[i] {
			num -= intArr[i]
			result += romanArr[i]
		}
	}
	return result
}
