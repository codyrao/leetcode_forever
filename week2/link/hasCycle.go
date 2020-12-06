package main

func main() {

}

func hasCycle(head *ListNode) bool {

	if nil == head || nil == head.Next{
		return false
	}

	curNode := head
	preNode := head

	for ; nil != curNode && nil != preNode && nil != preNode.Next; {
		curNode = curNode.Next
		preNode = preNode.Next.Next
		if curNode == preNode {
			return true
		}
	}

	return false
}
