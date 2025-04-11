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

// 反转k个节点的链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 创建一个哨兵节点
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for {

		// 检查剩余长度是否大于等于 k,每次获取k个
		count := 0
		temp := cur
		for temp != nil && count < k {
			temp = temp.Next
			count++
		}
		if count < k {
			break
		}

		// 反转 k 个节点
		var prev *ListNode
		tail := cur
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		// 连接反转后的部分
		pre.Next = prev
		tail.Next = cur

		// 更新 pre 位置，准备下一轮反转
		pre = tail
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
