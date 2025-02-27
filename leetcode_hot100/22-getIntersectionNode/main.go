package main

func main() {

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	h1 := headA
	h2 := headB
	for {
		for {
			if h2 != nil {
				if h2 == h1 {
					return h1
				}
				h2 = h2.Next
			} else {
				h2 = headB
				break
			}
		}
		if h1.Next != nil {
			h1 = h1.Next
		} else {
			return nil
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
