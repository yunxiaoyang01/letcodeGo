package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Score int
	Index int
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	var num1, num2 int
	num1, _ = strconv.Atoi(input.Text())
	input.Scan()
	num2, _ = strconv.Atoi(input.Text())
	arr := []*Student{}
	for i := 0; i < num1; i++ {
		input.Scan()
		stu := strings.Split(input.Text(), " ")
		score, _ := strconv.Atoi(stu[1])
		arr = append(arr, &Student{
			Name:  stu[0],
			Score: score,
			Index: i,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		if num2 == 0 {
			if arr[i].Score == arr[j].Score {
				return arr[i].Index < arr[j].Index
			} else {
				return arr[i].Score > arr[j].Score
			}
		} else {
			if arr[i].Score == arr[j].Score {
				return arr[i].Index < arr[j].Index
			} else {
				return arr[i].Score < arr[j].Score
			}
		}
	})
	for _, stu := range arr {
		fmt.Printf("%s %d\n", stu.Name, stu.Score)
	}
}
