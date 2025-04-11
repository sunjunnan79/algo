package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

// 统计出现频率最高的 K 个元素
func topKFrequent(nums []int, k int) []int {
	h := &minHeap{} // 这里要使用指针
	heap.Init(h)    // 初始化堆

	// 统计数字出现频率
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}

	// 维护小根堆
	for key, freq := range hash {
		heap.Push(h, KV{val: freq, key: key}) // 插入元素
		if h.Len() > k {
			heap.Pop(h) // 维持堆的大小为 k
		}
	}

	// 取出前 K 个元素
	res := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(KV).key) // 获取出现频率最高的数字
	}
	
	return res
}

// KV 结构体 (频率, 值)
type KV struct {
	val int // 频率
	key int // 元素值
}

// 实现最小堆
type minHeap []KV

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val } // 频率小的优先级高
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 插入元素到堆
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

// Pop 弹出元素（注意是指针方法）
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]  // 取出最后一个元素
	*h = old[:n-1] // 删除最后一个元素
	return x
}
