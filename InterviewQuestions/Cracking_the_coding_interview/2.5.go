/*
 * Given a circular linked list, implement an algorithm which returns node at the beginning
 * of the loop.
 *
 * DEFINITION
 *
 * Circular linked list: A (corrupt) linked list in which a node’s next pointer points to
 * an earlier node, so as to make a loop in the linked list.
 *
 * EXAMPLE
 *
 * Input: A -> B -> C -> D -> E -> C [the same C as earlier]
 *
 * Output: C
 */
package main
import (
	"fmt"
)

type Node struct {
	next *Node
	value string
}

// http://hawstein.com/posts/2.5.html
// http://umairsaeed.com/2011/06/23/finding-the-start-of-a-loop-in-a-circular-linked-list/
/***************************************************
http://nghiaho.com/?p=2063 (better explanation)

 1. detect circular linked list with Tortoise and Hare Algorithm
 2. find start of cycle

Algorithm
A = # nodes before cycle
B = # steps in cycle
L = # total nodes in cycle

A + modulus(B, L) = C    (definition 1)
A + B = L    (definition 2)

When the turtle reached node C, it has traveled A+B distance.
The rabbit also reached C by traveling 2(A+B) distance.
The rabbit traveled an extra (A+B) and still came back to the same node C.
This implies the loop cycle length is (A+B).

---
Advance A nodes on both sides
A + modulus(B + A, L) = (C advanced A nodes)

Substituting definition 2 for B + A gives

A + modulus(L, L) = (C advanced A nodes)
A + 0 = (C advanced A nodes)
A = (C advanced A nodes)
*****************************************************/

func findCycleStart(n *Node) *Node {
	slow, fast := n, n
	for fast != nil && fast.next != nil{
		if (fast.next != nil) {
			fast = fast.next.next
		}
		slow = slow.next
		if (fast == slow) {
			break
		}
	}
	if (fast == nil || fast.next == nil) { return nil }
  slow = n
	for fast != slow {
		slow = slow.next
		fast = fast.next
	}
	return fast
}

func construct(strArray []string, nth int) (*Node, *Node) {
	var n, head *Node
	for _, str := range strArray {
		if (n == nil) {
			n = &Node{value: str}
			head = n
		} else {
			n.next = &Node{value: str}
			n = n.next
		}
	}
	for e, i := head, 0; i < nth; e, i = e.next, i+1 {
		n.next = e
	}
	return head, n.next
}

func main() {
	head, start := construct([]string{"a", "b", "c", "d", "e", "f"}, 3)
	found := findCycleStart(head)
	fmt.Println(start == found)
	fmt.Println(start == head)
}
