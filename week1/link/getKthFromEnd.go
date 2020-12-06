/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package link

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if nil == head {
		return nil
	}
	curNode := head
	length := 0
	for ; nil != curNode; {
		curNode = curNode.Next
		length++
	}
	num := length - k
	if num < 0 {
		return nil
	}
	curNode = head
	for ; num != 0; {
		curNode = curNode.Next
		num--
	}
	return curNode
}
