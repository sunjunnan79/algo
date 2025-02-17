package main

import "fmt"

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(isPalindrome(h))

}

func isPalindrome(head *ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
