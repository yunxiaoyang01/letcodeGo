package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {

	//println(myAtoiSecond("-+1"))
	println(myAoti("-+1"))
}
func myAtoiFisrt(str string) int {
	//sign 表示符号
	//x 表示整数
	sign, x := 1, 0
	strTrim := strings.TrimLeft(str, " ")
	if strTrim == "" {
		return 0
	}
	strFirst := strTrim[0]

	if strFirst == 45 { //如果第一位 是（'-'） sign =-1
		sign = -1
	} else if strFirst == 43 { //如果第一位 是（'+'） sign =1
		sign = 1
	} else if strFirst > 47 && strFirst < 58 { //如果第一位是数字 sign=1 并且将第一位加入到整数中
		sign = 1
		x = int(strFirst - 48)
	} else {
		return 0
	}
	for i := 1; i < len(strTrim); i++ {
		if strTrim[i] > 47 && strTrim[i] < 58 {
			if x > math.MaxInt32/10 || (x == math.MaxInt32/10 && (int(strTrim[i])-48) > 7) { //因为x没有符号只需要判断是否大于int最大值
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			x = x*10 + (int(strTrim[i]) - 48)
		} else {
			break
		}
	}
	return sign * x
}

func myAtoiSecond(str string) int {
	// 第一步 去掉无效字符
	str = strings.TrimSpace(str)
	if str == "" || (len(str) == 1 && (str < "0" || str > "9")) {
		return 0
	}
	//校验首字符
	flag := ""
	if string(str[0]) == "-" {
		flag = "-"
		str = str[1:len(str)]
	} else if string(str[0]) == "+" {
		str = str[1:len(str)]
	}
	resStr := "0"
	//处理剩余字符
	for i := 0; i < len(str); i++ {
		if string(str[i]) < "0" || string(str[i]) > "9" {
			break
		}
		resStr += string(str[i])
	}
	resStr = flag + resStr
	res, err := strconv.ParseInt(resStr, 10, 32)
	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	//step5：转换异常处理
	if err != nil {
		if flag == "" {
			return MaxInt
		}
		return MinInt
	}
	return int(res)
}
func myAoti(str string) int {
	//step1：去无效字符
	str = strings.TrimSpace(str)
	if str == "" || (len(str) == 1 && (str < "0" || str > "9")) {
		return 0
	}
	//step2：规范首字符
	flag := ""
	if string(str[0]) == "-" {
		flag = "-"
		str = str[1:len(str)]
	} else if string(str[0]) == "+" {
		str = str[1:len(str)]
	}
	//step3：遍历检测数字0-9
	resStr := "0"
	for i := 0; i < len(str); i++ {
		if string(str[i]) < "0" || string(str[i]) > "9" {
			break
		}
		resStr += string(str[i])
	}
	resStr = flag + resStr
	//step4：转换
	res, err := strconv.ParseInt(resStr, 10, 32)

	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	//step5：转换异常处理
	if err != nil {
		if flag == "" {
			return MaxInt
		}
		return MinInt
	}

	return int(res)
}
