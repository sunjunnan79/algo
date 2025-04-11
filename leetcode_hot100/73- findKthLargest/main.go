package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Len() > k {
			h.Pop()
		}
	}

	return h.Pop()
}

type Heap struct {
	heap []int
}

func (h *Heap) Len() int { return len(h.heap) }

func (h *Heap) Pop() int {
	n := h.heap[0]
	h.heap = h.heap[1:h.Len()]
	return n
}

func (h *Heap) Push(x int) {
	if len(h.heap) == 0 {
		h.heap = append(h.heap, x)
		return
	}

	var n = []int{x}
	var left, right = 0, h.Len() - 1
	for left < right {
		mid := (left + right) / 2
		switch {
		case h.heap[mid] > x:
			right = mid - 1
		case h.heap[mid] < x:
			left = mid + 1
		default:
			h.heap = append(h.heap[:mid], append(n, h.heap[mid:]...)...)
			return
		}
	}

	if x > h.heap[left] {
		h.heap = append(h.heap[:left+1], append(n, h.heap[left+1:]...)...)
	} else {
		h.heap = append(h.heap[:left], append(n, h.heap[left:]...)...)
	}
}
