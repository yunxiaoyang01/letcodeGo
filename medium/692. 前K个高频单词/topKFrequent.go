package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	fmt.Printf("%s", topKFrequent(words, k))
}

func topKFrequent(words []string, k int) []string {
    cnt := map[string]int{}
    for _, w := range words {
        cnt[w]++
    }
    uniqueWords := make([]string, 0, len(cnt))
    for w := range cnt {
        uniqueWords = append(uniqueWords, w)
    }
    sort.Slice(uniqueWords, func(i, j int) bool {
        s, t := uniqueWords[i], uniqueWords[j]
        return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
    })
    return uniqueWords[:k]
}