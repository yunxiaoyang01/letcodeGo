package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var line1 int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line1, _ = strconv.Atoi(input.Text())
	var line2 string
	input.Scan()
	line2 = input.Text()
	arr := strings.Split(line2, " ")
	if line1 != len(arr) {
		arr = arr[0:line1]
	}
	var line3 int
	input.Scan()
	line3, _ = strconv.Atoi(input.Text())

	res := ""
	if line3 == 0 {
		sort.Slice(arr, func(i, j int) bool {
			first, _ := strconv.Atoi(arr[i])
			second, _ := strconv.Atoi(arr[j])
			return first < second
		})

	} else {
		sort.Slice(arr, func(i, j int) bool {
			first, _ := strconv.Atoi(arr[i])
			second, _ := strconv.Atoi(arr[j])
			return first > second
		})

	}
	res = strings.Join(arr, " ")
	fmt.Printf("%s\n", res)
}
