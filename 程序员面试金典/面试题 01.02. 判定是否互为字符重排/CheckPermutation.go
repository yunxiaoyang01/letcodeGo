package main

func main() {

}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1, m2 := map[int]int{}, map[int]int{}
	for i := range s1 {
		m1[int(s1[i]-'a')] += 1
	}
	for i := range s2 {
		m2[int(s2[i]-'a')] += 1
	}
	for key, value := range m1 {
		if m2[key] != value {
			return false
		}
	}
	return true
}
