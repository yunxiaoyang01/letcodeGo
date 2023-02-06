package main

import "fmt"

func main() {
	var num int
	for {
		_, err := fmt.Scanf("0x%x", &num)
		if err != nil {
			return
		}
		fmt.Println(num)
	}
}
