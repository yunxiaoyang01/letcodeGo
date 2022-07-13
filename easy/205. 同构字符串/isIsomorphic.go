package main

import "fmt"

func main() {
	fmt.Printf("%v", isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if sMap[x] > 0 && sMap[x] != y || tMap[y] > 0 && tMap[y] != x {
			return false
		}
		sMap[x] = y
		tMap[y] = x
	}
	return true
}
