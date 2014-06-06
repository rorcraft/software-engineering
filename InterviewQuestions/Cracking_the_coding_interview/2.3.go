/*
 * Implement an algorithm to delete a node in the middle of a single linked list,
 * given only access to that node.
 * EXAMPLE
 * Input: the node ‘c’ from the linked list a->b->c->d->e Result: nothing is returned,
 * but the new linked list looks like a->b->d->e
 */
package main
import (
	"fmt"
)

type Node struct {
	next *Node
	value string
}

func deleteNode(node *Node) {
	// copy value
	if (node.next != nil) {
		node.value = node.next.value
		node.next = node.next.next
	}
}
func main() {
	var middle, head, n *Node
	for _, str := range []string{"a", "b", "c", "d", "c", "b"} {
		if (n == nil) {
			n = &Node{value: str}
			head = n
		} else {
			n.next = &Node{value: str}
			n = n.next
		}

		if str == "d" {
			middle = n
		}
	}
	deleteNode(middle)
	n = head
	for n != nil {
		fmt.Printf("%s ", n.value)
		n = n.next
	}
	fmt.Println()
}
