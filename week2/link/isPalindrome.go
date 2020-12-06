package main

func main() {

}

var leftNode *ListNode

func isPalindrome(head *ListNode) bool {
	if nil == head {
		return true
	}
	tmp := make([]int, 0)
	curNode := head
	for ; nil != curNode; {
		tmp = append(tmp, curNode.Val)
		curNode = curNode.Next
	}
	left := 0
	right := len(tmp) - 1

	for ; left < right; {
		if tmp[left] != tmp[right] {
			return false
		}
		left++
		right--
	}

	return true

}
