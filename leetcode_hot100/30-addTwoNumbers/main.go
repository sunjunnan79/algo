package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resp []int
	var m int
	var ad int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			m += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m += l2.Val
			l2 = l2.Next
		}

		resp = append(resp, (ad+m)%10)
		ad = (ad + m) / 10
		m = 0
	}
	if ad > 0 {
		resp = append(resp, ad)
	}

	var head = &ListNode{}
	var h = head
	n := len(resp)
	for i := 0; i < n; i++ {
		head.Next = &ListNode{Val: resp[i]}
		head = head.Next
	}
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
