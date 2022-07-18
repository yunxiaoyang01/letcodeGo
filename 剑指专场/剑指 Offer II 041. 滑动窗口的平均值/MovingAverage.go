package main

import "fmt"

type MovingAverage struct {
	Size int
	Arr  []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size: size,
		Arr:  []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	var length int
	if len(this.Arr) < this.Size {
		this.Arr = append(this.Arr, val)
		length = len(this.Arr)
	} else if len(this.Arr) == this.Size {
		this.Arr = append(this.Arr, val)
		this.Arr = this.Arr[1:]
		length = this.Size
	} else {
		this.Arr = this.Arr[this.Size:]
		length = this.Size
	}
	sum := float64(0)
	for _, num := range this.Arr {
		sum += float64(num)
	}
	return sum / float64(length)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {
	m := Constructor(3)
	fmt.Println(m.Next(1))
	fmt.Println(m.Next(10))
	fmt.Println(m.Next(3))
	fmt.Println(m.Next(5))
}
