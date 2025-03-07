package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: -4,
				},
			},
		},
	}
	head.Next.Next.Next = head
	fmt.Println(detectCycle(head))
}

func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head
	for {
		//检查是否到头
		if fast == nil || fast.Next == nil || slow == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 找到环的起始节点
			// 让 ptr 和 slow 一起走，每次走一步，直到它们相遇
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			// 返回环的起始节点
			return head
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
