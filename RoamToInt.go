package main

import "strings"

func main() {
	println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	map1 := map[string]int{ //正常情况
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	map2 := map[string]int{ //特殊情况
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var num []int
	var number int
	for key, value := range map2 { //先枚举特殊情况，把特殊情况转化成数字然后把特殊罗马字符变成""
		if strings.Contains(s, key) {
			showCount := strings.Count(s, key)
			for i := 0; i < showCount; i++ {
				num = append(num, value)
			}
			s = strings.Replace(s, key, "", showCount)
		}
	}
	for key, value := range map1 { //最后枚举正常把包含的数字加入到num数组中
		if strings.Contains(s, key) {
			showCount := strings.Count(s, key)
			for i := 0; i < showCount; i++ {
				num = append(num, value)
			}
			s = strings.Replace(s, key, "", showCount)
		}
	}
	for i := 0; i < len(num); i++ { //加和
		number += num[i]
	}
	return number
}
