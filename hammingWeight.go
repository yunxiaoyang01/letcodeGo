package main

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 { //判断尾数是不1
			count++
		}
		num >>= 1 //数字右移一位
	}
	return count
}
