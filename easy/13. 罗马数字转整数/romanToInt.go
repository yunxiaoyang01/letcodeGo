package main

import "fmt"

func main() {
	fmt.Printf("%d", romanToInt("XL"))
}
func romanToInt(s string) int {
	var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			res -= value
		} else {
			res += value
		}
	}
	return res
}
