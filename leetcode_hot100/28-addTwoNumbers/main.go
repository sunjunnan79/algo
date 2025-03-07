package main

func main() {
	h1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	h2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(h1, h2)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resp []int
	var m int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			m += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m += l2.Val
			l2 = l2.Next
		}
		resp = append(resp, m)
		m = 0
	}
	for i:=0; i<len(resp); i++ {
		if resp[n]
	}
	//操作并添加到最后
	var head = &ListNode{}
	var h = head
	n := len(resp)
	for i := 0; i < n; i++ {
		head.Next = &ListNode{Val: resp[n-i-1]}
		head = head.Next
	}
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
