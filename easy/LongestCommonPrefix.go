package easy

import (
	"math"
	"strings"
)

//func main() {
//	var strs []string
//	strs = append(strs, "dog")
//	strs = append(strs, "racecar")
//	strs = append(strs, "car")
//	println(longestCommonPrefix(strs))
//}
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	var minlen int = math.MaxInt32
	for _, str := range strs {
		if len(str) < minlen {
			minlen = len(str)
		}
	}
	prefix := strs[0][0:minlen]
	for {
		count := 0
		for i := 0; i < n; i++ {
			if strings.HasPrefix(strs[i], prefix) {
				count += 1
			}
		}
		if count == n {
			break
		} else {
			prefix = prefix[0 : len(prefix)-1]
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}
