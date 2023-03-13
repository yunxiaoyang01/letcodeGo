package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("%d", minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}))
	fmt.Printf("%d", minNumberOfHours(2, 4, []int{1}, []int{3}))
}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	sumEnergy := 0
	for _, v := range energy {
		sumEnergy += v
	}
	needEnegy := 0
	if sumEnergy >= initialEnergy {
		needEnegy = sumEnergy - initialEnergy + 1
	}
	needExp := 0
	tmpExp := initialExperience
	for _, v := range experience {
		if v < tmpExp {
			tmpExp += v
		} else {
			needExp += v - tmpExp + 1
			tmpExp = 2*v + 1
		}
	}
	return needEnegy + needExp
}
