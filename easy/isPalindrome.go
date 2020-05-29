package easy

import (
	"strings"
	"unicode"
)

//func main() {
//	println(isPalindrome1251("A man, a plan, a canal: Panama"))
//}
func isPalindrome125(s string) bool {

	s = strings.ToLower(s) // 字母转换为小写
	l, r := 0, len(s)-1
	for l < r {
		if !('a' <= s[l] && s[l] <= 'z' || '0' <= s[l] && s[l] <= '9') { // 只处理字母和数字
			l++
			continue
		}
		if !('a' <= s[r] && s[r] <= 'z' || '0' <= s[r] && s[r] <= '9') {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
func isPalindrome1251(s string) bool {
	a := strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	s2 := strings.Join(a, "")
	for i := range s2 {
		if s2[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}
