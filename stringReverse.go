package main

import (
	"strconv"
	"strings"
)

func main() {
	var x = -2147483648
	println(reverse(x))

}

func reverse(x int) int {
	if x >= 0 && x < 10 {
		return x
	}
	xStr := strconv.Itoa(x)         //x转换成字符串
	var INT_MIN int64 = -2147483648 //定义int最小值
	var INT_MAX int64 = 2147483647  //定义int最大值
	var y int64 = 0
	if strings.HasPrefix(xStr, "-") { //判断整数前缀是否是负数
		xNew := strings.Replace(xStr, "-", "", 1) //去掉负号
		xDao := reverseString(xNew)               //反转新的整数
		xDao = "-" + xDao
		y, _ = strconv.ParseInt(xDao, 10, 64)
		if y < INT_MIN || y > INT_MAX { //判断是否在int范围之内
			return 0
		}
	} else {
		xDao := reverseString(xStr)
		y, _ = strconv.ParseInt(xDao, 10, 64)
		if y < INT_MIN || y > INT_MAX {
			return 0
		}
	}
	yStr := strconv.FormatInt(y, 10) //字符串转整数
	result, _ := strconv.Atoi(yStr)
	return result

}
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
