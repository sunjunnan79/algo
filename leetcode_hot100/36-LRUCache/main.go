package main

import "fmt"

func main() {
	lRUCache := Constructor(1)
	fmt.Println(lRUCache.Get(6)) // 返回 1
	fmt.Println(lRUCache.Get(8)) // 返回 3 (未找到)
	lRUCache.Put(12, 1)
	fmt.Println(lRUCache.Get(2))
	lRUCache.Put(15, 11)
	lRUCache.Put(5, 2)
	lRUCache.Put(1, 15)
	lRUCache.Put(4, 2)
	fmt.Println(lRUCache.Get(4)) // 返回 1
	lRUCache.Put(15, 15)
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
	_, ok := this.Cache[key]
	if ok {
		//查找目标节点并取出
		h := &ListNode{
			Next: this.Head,
		}
		for h.Next != nil {
			if h.Next.Val == key {
				//添加到尾结点
				this.Tail.Next = &ListNode{Val: key}
				this.Tail = this.Tail.Next
				//删除当前节点
				if this.Head == h.Next {
					this.Head = h.Next.Next
				} else {
					h.Next = h.Next.Next
				}

				break
			}
			h = h.Next
		}
		return this.Cache[key]
	} else {
		return -1
	}

}

func (this *LRUCache) Put(key int, value int) {
	//查是否存在
	_, ok := this.Cache[key]
	if !ok {
		//如果不存在且容量已经满了,需要进行缓存淘汰
		if this.Size == this.Length {

			//删除最老的值
			delete(this.Cache, this.Head.Val)

			//移动头结点到下一个节点
			this.Head = this.Head.Next
			if this.Head == nil {
				this.Tail = this.Tail.Next
			}

		} else {
			this.Length++
		}

		//如果是未初始化的情况
		if this.Tail == nil {
			this.Tail = &ListNode{Val: key}
			this.Head = this.Tail
		} else {
			//记录到尾结点
			this.Tail.Next = &ListNode{Val: key}
			this.Tail = this.Tail.Next
		}
	} else {
		//查找目标节点并取出
		h := &ListNode{
			Next: this.Head,
		}
		for h.Next != nil {
			if h.Next.Val == key {
				//添加到尾结点
				this.Tail.Next = &ListNode{Val: key}
				this.Tail = this.Tail.Next
				//删除当前节点
				if this.Head == h.Next {
					this.Head = h.Next.Next
				} else {
					h.Next = h.Next.Next
				}

				break
			}
			h = h.Next
		}

	}
	//添加到缓存
	this.Cache[key] = value
}

type ListNode struct {
	Val  int
	Next *ListNode
}
