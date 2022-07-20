package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {
	a, b := 0, 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}
	return b
}
