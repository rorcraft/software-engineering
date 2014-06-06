/*
 * Implement an algorithm to find the nth to last element of a singly linked list.
 */

package main
import (
	"fmt"
)

type Node struct {
	next *Node
	value string
}

func nthLast(n int, node *Node) *Node {
	nth := node
	for e := node; e != nil; e = e.next {
		if (n >= 0) {
			n--
		} else {
			nth = nth.next
		}
	}
	if (n > 0) {
		// error
		panic("nth is out of range")
	}
	return nth
}

// avoid branching in loop
func nthLast2(n int, node *Node) *Node {
	var ahead *Node
	nth := node
	for ahead = node; ahead != nil && n >= 0; ahead = ahead.next {
		n--
	}
	if (n > 0) {
		// error
		panic("nth is out of range")
	}
	for ahead != nil {
		nth = nth.next
		ahead = ahead.next
	}
	return nth
}

func main() {
	var n, head *Node
	for _, str := range []string{"a", "b", "c", "d", "c", "b"} {
		if (n == nil) {
			n = &Node{value: str}
			head = n
		} else {
			n.next = &Node{value: str}
			n = n.next
		}
	}
	fmt.Println("nth:", 2, "from last is", nthLast(2, head).value)
	fmt.Println("nth:", 2, "from last is", nthLast2(2, head).value)
	fmt.Println("nth:", 4, "from last is", nthLast(4, head).value)
	fmt.Println("nth:", 4, "from last is", nthLast2(4, head).value)
}
