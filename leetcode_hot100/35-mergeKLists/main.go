package main

import "sort"

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	//转换成数组
	var l []int
	for _, list := range lists {
		for list != nil {
			l = append(l, list.Val)
			list = list.Next
		}
	}
	//排序
	sort.Ints(l)
	head := &ListNode{}
	h := head
	//二次生成链表
	for i := 0; i < len(l); i++ {
		head.Next = &ListNode{
			Val: l[i],
		}
		head = head.Next
	}
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
