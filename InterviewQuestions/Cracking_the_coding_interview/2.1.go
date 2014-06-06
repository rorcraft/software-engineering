/* Write code to remove duplicates from an unsorted linked list.
 * FOLLOW UP
 * How would you solve this problem if a temporary buffer is not allowed? O(n^2)
 *
 * http://hawstein.com/posts/2.1.html
 */
package main
import (
	"fmt"
)

type Node struct {
	next *Node
	value string
}

func removeDuplicate(n *Node) *Node {
	head := n
	check := make(map[string]bool)
	var prev *Node
	for e := n; e != nil; e = e.next {
		if (check[e.value]) {
			if (prev != nil) {
				prev.next = e.next
			}
		} else {
			check[e.value] = true
			prev = e
		}
	}
	return head
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
	d := removeDuplicate(head)
	for d != nil {
		fmt.Printf("%s ", d.value)
		d = d.next
	}
	fmt.Println()
}
