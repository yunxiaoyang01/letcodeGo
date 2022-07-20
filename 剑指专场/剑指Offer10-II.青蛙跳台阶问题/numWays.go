package main

func numWays(n int) int {
	a, b := 1, 2
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	for i := 3; i <= n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}
	return b
}
