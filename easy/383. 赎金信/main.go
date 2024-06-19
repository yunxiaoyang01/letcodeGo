package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
func canConstruct(ransomNote string, magazine string) bool {
	magzineMap := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		magzineMap[magazine[i]]++
	}
	ransomNoteMap := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		ransomNoteMap[ransomNote[i]]++
	}
	for k, v := range ransomNoteMap {
		if magzineMap[k] < v {
			return false
		}
	}
	return true
}
