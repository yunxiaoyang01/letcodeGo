package easy

func main() {
	println(titleToNumber("ABA"))
}

func titleToNumber(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		mi := int(s[i] - 'A' + 1)
		sum = sum*26 + mi
	}
	return sum
}
