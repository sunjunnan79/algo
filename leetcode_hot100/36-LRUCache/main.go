package main

func main() {
	obj := Constructor(10)
	param_1 := obj.Get(key)
	obj.Put(key, value)
}

type LRUCache struct {
	Cache  map[int]int
	Head   *ListNode
	Tail   *ListNode
	Size   int
	Length int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:  make(map[int]int, capacity),
		Size:   capacity,
		Length: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	//移除当前使用的节点
	RemoveNodeFromHead(this.Head, key)
	//在尾结点添加
	return this.Cache[key]
}

func (this *LRUCache) Put(key int, value int) {
	//查是否存在
	_, ok := this.Cache[key]
	if !ok {
		if this.Size == this.Length {
			//去掉最早插入的值
			this.RemoveNodeFromHead(key)
			//添加到队列末尾
			this.AddNodeToTail(key)
		}
	}

	//添加到缓存
	this.Cache[key] = value
	//尾结点添加
	this.Tail.Next = &ListNode{Val: key}
}

// 移除指定节点
func (this *LRUCache) RemoveNodeFromHead(key int) {
	h := &ListNode{
		Next: this.Head,
	}
	for h.Next != nil {
		if h.Next.Val == key {
			h.Next = h.Next.Next
		}
		h = h.Next
	}
	return
}

// 添加到尾结点
func (this *LRUCache) AddNodeToTail(v int) {
	if this.Tail == nil {
		this.Head
	}
	this.Tail.Next = &ListNode{
		Val: v,
	}
	this.Tail = this.Tail.Next
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}
