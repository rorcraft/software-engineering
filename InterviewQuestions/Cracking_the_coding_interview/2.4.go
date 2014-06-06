/*
 * You have two numbers represented by a linked list, where each node contains a single digit.
 * The digits are stored in reverse order, such that the 1’s digit is at the head of the list.
 * Write a function that adds the two numbers and returns the sum as a linked list.
 *
 * EXAMPLE
 *
 * Input: (3 -> 1 -> 5), (5 -> 9 -> 2)
 *
 * Output: 8 -> 0 -> 8
 */
package main
import  (
	"fmt"
)

type Node struct {
	next *Node
	value int
}

func add(one, another *Node) *Node {
	carry := 0
	var pointer, resultHead *Node
	if (one == nil && another == nil) {
		return nil
	} else if one == nil {
		return another
	} else if another == nil {
		return one
	}
	for one != nil || another != nil || carry > 0 {
		result := &Node{}
		if (one != nil && another != nil) {
			sum := one.value + another.value + carry
			carry = sum / 10
			result.value = sum % 10
			one = one.next
			another = another.next
		} else if (one != nil) {
			result.value = one.value
			one = one.next
		} else if (another != nil) {
			result.value = another.value
			another = another.next
		} else if (carry > 0) {
			result.value = carry
			carry = 0
		}
		if (pointer == nil) {
			resultHead = result
			pointer = result
		} else {
			pointer.next = result
			pointer = pointer.next
		}
	}
	return resultHead
}

func construct(strArray []int) *Node {
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
	return head
}

func printNode(n *Node) {
	for e := n; e != nil; e = e.next {
		fmt.Print(e.value)
		if (e.next != nil) {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func main() {
	one := construct([]int{5, 5, 5})
	another := construct([]int{5, 5, 5})
	sum := add(one, another)

	printNode(one)
	printNode(another)
	fmt.Print("Sum: ")
	printNode(sum)


	one = construct([]int{5, 1, 9})
	another = construct([]int{5, 1, 9})
	sum = add(one, another)

	printNode(one)
	printNode(another)
	fmt.Print("Sum: ")
	printNode(sum)
}
