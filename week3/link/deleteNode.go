package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {

	if nil == node || nil == node.Next {
		return
	}

	*node = *node.Next

}
