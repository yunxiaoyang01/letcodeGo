package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var num int
	input.Scan()
	num, _ = strconv.Atoi(input.Text())
	arr := []string{}
	for i := 0; i < num; i++ {
		input.Scan()
		arr = append(arr, input.Text())
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for _, v := range arr {
		fmt.Printf("%s\n", v)
	}
}
