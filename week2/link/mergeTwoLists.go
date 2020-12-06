/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if nil == l1 {
		return l2
	}

	if nil == l2 {
		return l1
	}

	preNode := new(ListNode)
	resNode := preNode
	for ; nil != l1 && nil != l2; {

		if l1.Val <= l2.Val {
			resNode.Next = l1
			l1 = l1.Next
			resNode = resNode.Next
			continue
		}

		resNode.Next = l2
		l2 = l2.Next
		resNode = resNode.Next
	}
	if nil != l1 {
		resNode.Next = l1
	}
	if nil != l2 {
		resNode.Next = l2
	}

	return preNode.Next
}
