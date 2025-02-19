package main

import "fmt"

func main() {
	//list1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//list2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	var list1 *ListNode
	list2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l := mergeTwoLists(list1, list2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode = nil
	//初始化两个值
	var v1, v2 *int

	if list1 != nil {
		v1 = &list1.Val
	}

	if list2 != nil {
		v2 = &list2.Val
	}

	//根据两个值的大小进行连败
	switch {
	case v1 == nil && v2 == nil:
		return head
	case v1 == nil || v2 != nil && *v1 >= *v2:
		head = list2
		list2 = list2.Next
	case v2 == nil || v1 != nil && *v2 >= *v1:
		head = list1
		list1 = list1.Next
	}
	var h = head

	// 正式开始遍历
	for {
		//初始化两个值
		var v1, v2 *int

		if list1 != nil {
			v1 = &list1.Val
		}

		if list2 != nil {
			v2 = &list2.Val
		}

		//根据两个值的大小进行连败
		switch {
		case v1 == nil && v2 == nil:
			return h
		case v1 == nil || v2 != nil && *v1 >= *v2:
			head.Next = list2
			head = head.Next
			list2 = list2.Next
		case v2 == nil || v1 != nil && *v2 >= *v1:
			head.Next = list1
			head = head.Next
			list1 = list1.Next
		}

	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
