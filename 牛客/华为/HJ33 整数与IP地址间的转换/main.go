package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		arr := strings.Split(str, ".")
		if len(arr) == 1 {
			ipInt, _ := strconv.Atoi(str)
			fmt.Printf("%s\n", IntoIp(ipInt))
		} else {
			fmt.Printf("%d\n", IptoIn(str))

		}

	}
}

func IptoIn(str string) int {
	arr := strings.Split(str, ".")
	ans := 0
	for _, ip := range arr {
		ipInt, _ := strconv.Atoi(ip)
		ans = ans*256 + ipInt
	}
	return ans
}

func IntoIp(num int) string {
	i := 1
	temp := 0
	nums := []int{}
	for num > 0 {
		temp += (num % 2) * i
		num /= 2
		i *= 2
		if i == 256 {
			nums = append(nums, temp)
			i = 1
			temp = 0
		}
	}
	if temp != 0 {
		nums = append(nums, temp)
	}
	if len(nums) < 4 {
		nums = append(nums, 0)
	}
	s := ""
	for _, v := range nums {
		s = fmt.Sprintf("%s.%s", strconv.Itoa(v), s)
	}
	return s[:len(s)-1]
}
