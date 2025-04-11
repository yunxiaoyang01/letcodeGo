package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor()
	rs.Insert(0)
	rs.Insert(1)
	rs.Remove(0)
	rs.Insert(2)
	rs.Remove(1)
	rs.GetRandom()
	fmt.Printf("%+v", rs.arr)
}

type RandomizedSet struct {
	arr []int
	m   map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.m[val]
	if !ok {
		return false
	}
	last := len(this.arr) - 1
	this.arr[id] = this.arr[last]
	this.m[this.arr[id]] = id
	this.arr = this.arr[:last]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
