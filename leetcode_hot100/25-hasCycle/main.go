package main

func main() {

}

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)

	for {

		if head != nil {
			if !hash[head] {
				hash[head] = true
			} else {
				return true
			}

		} else {
			return false
		}
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
