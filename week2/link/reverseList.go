/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	preNode := head
	var curNode *ListNode = nil
	var tmp *ListNode
	for ; preNode != nil; {
		tmp = preNode.Next
		preNode.Next = curNode
		curNode = preNode
		preNode = tmp

	}

	return curNode
}
