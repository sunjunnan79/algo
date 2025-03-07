package main

import "sort"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//转换链表为数组
	var l []int
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	//排序
	sort.Ints(l)
	head = &ListNode{}
	h := head
	//排序之后重新构建链表
	for i := 0; i < len(l); i++ {
		head.Next = &ListNode{
			Val: l[i],
		}
		head = head.Next
	}
	return h.Next
}
