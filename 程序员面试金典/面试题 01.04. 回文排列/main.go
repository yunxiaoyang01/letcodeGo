package main

func main() {

}

func canPermutePalindrome(s string) bool {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	count := 0
	for _, v := range m {
		if v%2 != 0 {
			count++
		}
	}
	if count != 1 && count != 0 {
		return false
	}
	return true
}
