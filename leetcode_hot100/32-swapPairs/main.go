package main

func main() {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	swapPairs(head)
}

func swapPairs(head *ListNode) *ListNode {
	// 哨兵节点，方便统一返回
	var h = &ListNode{}
	h.Next = head
	cur := h

	for cur.Next != nil && cur.Next.Next != nil {
		a := cur.Next      // 第一个节点
		b := cur.Next.Next // 第二个节点
		c := b.Next        // 下一组的头节点

		// 交换
		cur.Next = b
		b.Next = a
		a.Next = c

		// 移动指针到下一组
		cur = a
	}

	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
