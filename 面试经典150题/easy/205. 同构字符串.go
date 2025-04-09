package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		x, y := s[i], t[i]
		if sMap[x] > 0 && sMap[x] != y || tMap[y] > 0 && tMap[y] != x {
			return false
		}
		sMap[x] = y
		tMap[y] = x
	}
	return true
}
