package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	result := reverseListV2(head)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseListV2(head *ListNode) *ListNode {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	ret := &ListNode{}
	curr := ret
	for i := len(result) - 1; i >= 0; i-- {
		curr.Next = &ListNode{
			Val: result[i],
		}
		curr = curr.Next
	}
	return ret.Next
}
