package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	rMap := map[byte]int{}
	mMap := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		rMap[ransomNote[i]]++
	}
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]]++
	}
	for k, v := range rMap {
		if mMap[k] < v {
			return false
		}
	}
	return true
}
