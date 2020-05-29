package easy

import (
	"strconv"
	"strings"
)

//func main() {
//	println(countAndSay(3))
//}

func countAndSay(n int) string {
	// 递归出口
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)
	var b strings.Builder
	cnt, cur := 1, pre[0]
	for i := range pre[1:] {
		if pre[i+1] == cur {
			cnt++
		} else {
			// 将 当前数字的数目 和 当前数字 追加到 Builder
			b.WriteString(strconv.Itoa(cnt))
			b.WriteByte(cur)
			cur = pre[i+1]
			cnt = 1
		}
	}
	b.WriteString(strconv.Itoa(cnt))
	b.WriteByte(cur)
	return b.String()
}
