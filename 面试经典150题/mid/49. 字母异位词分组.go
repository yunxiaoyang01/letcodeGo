package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortted := string(s)
		m[sortted] = append(m[sortted], str)
	}
	ans := make([][]string,0)
	for _,v := range m {
		ans = append(ans, v)
	}
	return ans
}
