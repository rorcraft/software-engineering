### Floyd's Cycle-Finding Algorithm or Tortoise and Hare Algorithm

http://codingfreak.blogspot.com/2012/09/detecting-loop-in-singly-linked-list_22.html

2 pointers namely slow Pointer and fast Pointer to traverse a Singly Linked List at different speeds. A slow Pointer (Also called Tortoise) moves one step forward while fast Pointer (Also called Hare) moves 2 steps forward

http://nghiaho.com/?p=2063 
 
__find start of cycle__

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

```go
type Node struct {
	next *Node
	value string
}

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
```

### Sorting

* use heapsort.
