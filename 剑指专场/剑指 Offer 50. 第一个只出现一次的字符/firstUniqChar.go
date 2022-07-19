package main

func main() {

}

func firstUniqChar(s string) byte {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
