/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

type Node struct {
	next *Node
	data interface{}
}

func main() {

}

func isCycle(node *Node) bool {

	p1 := node
	p2 := node

	for ; (nil != p2) && ( nil != p2.next); {
		p1 = p1.next
		p2 = p2.next.next

		if p1 == p2 {
			return true
		}

	}

	return false
}


