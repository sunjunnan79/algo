package main

func main() {
	removeNthFromEnd(&ListNode{Val: 1}, 1)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//建立头结点,进行操作
	h := &ListNode{}
	h.Next = head
	head = h
	var l []*ListNode
	for head != nil {
		l = append(l, head)
		head = head.Next
	}

	l[len(l)-n-1].Next = l[len(l)-n].Next
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
