/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if nil == node  {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
