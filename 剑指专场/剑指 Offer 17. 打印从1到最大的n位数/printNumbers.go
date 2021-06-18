package main

import "fmt"

func main() {
	fmt.Printf("%d",printNumbers(2))
}

func printNumbers(n int) []int {
	end := pow(n)-1
	res := []int{}
	for i:=1;i<=end;i++ {
		res = append(res,i)
	}
	return res
}

func pow(n int) int {
	sum := 1

	for i := 0; i < n; i++ {
		sum *= 10
	}
	return sum
}
