package main

import "fmt"

func main() {
	fmt.Println(jumpFloor(5))
}

func jumpFloor(number int) int {
	if number == 1 || number == 2 {
		return number
	}
	a, b := 1, 2
	for i := 3; i <= number; i++ {
		a, b = b, a+b
	}
	return b
}
