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
	res := getKthFromEndV2(head, k)
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

	for count != 0 {
		head = head.Next
		count--
	}
	return head
}

// 双指针走法
// 题解： 快慢指针 让快指针先走k步，然后快慢指针一起走，快指针走到头，slow指针走到的位置就是倒数第k个节点
func getKthFromEndV2(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
