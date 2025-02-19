package main

import "fmt"

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	h := head
	for i := 2; i <= 5; i++ {
		head.Next = &ListNode{
			Val: i,
		}
		head = head.Next
	}

	l := reverseKGroup(h, 2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var list []*ListNode
	var last = &ListNode{}
	var final = last
	var next = head

	for next != nil {
		list = append(list, next)
		next = next.Next
		if len(list) == k {

			//指向上一个
			for i := k - 1; i-1 >= 0; i-- {
				list[i].Next = list[i-1]
			}

			//上一次的末尾指针要指向当前的指针
			last.Next = list[k-1]
			list[0].Next = next

			//存储到上次的记录去
			last = list[0]

			//清除list
			list = list[:0]

		}

	}

	return final.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
