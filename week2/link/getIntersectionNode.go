package main

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if nil==headA  || nil == headB {
		return nil
	}


	curNode := headA
	lengthA := 0
	for ; curNode != nil; {
		lengthA++
		curNode = curNode.Next
	}

	curNode = headB
	lengthB := 0
	for ; curNode != nil; {
		lengthB++
		curNode = curNode.Next
	}
	var diff int
	if lengthA >= lengthB {
		diff = lengthA - lengthB
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}
	} else {
		diff = lengthB - lengthA
		for i := 0; i < diff; i++ {
			headB = headB.Next
		}
	}

	for ; headA != nil && headB != nil; {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
