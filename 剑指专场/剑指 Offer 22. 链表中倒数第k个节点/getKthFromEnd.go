package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	k := 2
	res := getKthFromEnd(head, k)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	sum := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		sum += 1
	}
	count := sum - k 

	for count > 0 {
		head = head.Next
		count--
	}
	return head
}
