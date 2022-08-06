package main

func main() {

}

func isUnique(astr string) bool {
	m := map[int]int{}
	for _, v := range astr {
		m[int(v-'a')] += 1
		if m[int(v-'a')] > 1 {
			return false
		}
	}
	return true
}

func isUniqueV2(astr string) bool {
	mark := 0
	for _, v := range astr {
		move := int(v - 'a')
		if (mark & (1 << move)) != 0 {
			return false
		} else {
			mark |= (1 << move)
		}
	}
	return true
}
