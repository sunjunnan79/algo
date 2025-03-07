package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	//获取node的map
	count := 0
	var nodeList []*Node
	l := make(map[*Node]int)
	for head != nil {
		l[head] = count
		nodeList = append(nodeList, &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		})
		count++
		head = head.Next
	}
	head = &Node{}
	h := head
	for i := 0; i < count; i++ {
		if nodeList[i].Random != nil {
			nodeList[i].Random = nodeList[l[nodeList[i].Random]]
		}

		head.Next = nodeList[i]
		head = head.Next
	}

	return h.Next
}
