package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 简单解法 顺序读出来然后反着写回去
func reverseList(head *ListNode) *ListNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := &ListNode{}
	temp := res
	for i := len(arr) - 1; i >= 0; i-- {
		temp.Next = &ListNode{
			Val: arr[i],
		}
		temp = temp.Next
	}
	return res.Next
}

// 指针迭代
func reverseListV2(head *ListNode) *ListNode {
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
