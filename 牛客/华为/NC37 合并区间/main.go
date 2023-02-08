package main

import "sort"

func main() {

}

type Interval struct {
	Start int
	End   int
}

func merge(intervals []*Interval) []*Interval {
	// write code here
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res := make([]*Interval, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= res[len(res)-1].End {
			res[len(res)-1].End = max(intervals[i].End, res[len(res)-1].End)
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
