package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}



// 第一种，按照题意一步一步遍历
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp.Next = l2
			break
		} else if l2 == nil {
			temp.Next = l1
			break
		} else if l1.Val <= l2.Val {
			temp.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			temp.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		temp = temp.Next
	}
	return res.Next
}

// 递归方式
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsV2(l1.Next,l2)
		return l1
	}else {
		l2.Next = mergeTwoListsV2(l1,l2.Next)
		return l2
	}
}