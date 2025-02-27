package main

import "fmt"

func main() {

	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	h := head
	for i := 1; i < 5; i++ {
		head.Next = &ListNode{
			Val: i,
		}
		head = head.Next
	}

	l := reverseList(h)
	for {

		if l != nil {
			fmt.Println(l.Val)
		} else {
			break
		}
		l = l.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var last *ListNode = nil
	var next = head

	for {

		next = next.Next
		if next != nil {
			// 修改当前节点的下一个节点为上一个节点
			head.Next = last

			// 保存当前节点为上一个节点
			last = head
			//移动到下一个节点
			head = next
		} else {
			head.Next = last
			// 保存当前节点为上一个节点
			last = head
			break
		}

	}

	return head
}
