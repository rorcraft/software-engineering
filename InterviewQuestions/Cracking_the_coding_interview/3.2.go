/*
 * How would you design a stack which, in addition to push and pop,
 * also has a function min which returns the minimum element?
 * Push, pop and min should all operate in O(1) time.
 */

// http://hawstein.com/posts/3.2.html

// storing the current min with the stack.
// min could be the address of the mem to save space.
package main
import (
	"fmt"
)

type Node struct {
	val int
	min int
}

type Stack struct {
	cur int
	mem []Node
}

func (s *Stack) alloc(size int) {
	s.cur = -1
	s.mem = make([]Node, size)
}

func (s *Stack) empty() bool {
	return s.cur == -1
}

func (s *Stack) push(num int) {
	if s.empty() {
		s.mem[0] = Node{val: num, min: num}
		s.cur = 0
	} else{
		node := s.mem[s.cur]
		next := Node{val: num}
		if node.min > num {
			next.min = num
		} else {
			next.min = node.min
		}
		s.cur++
		s.mem[s.cur] = next
	}
}

func (s *Stack) min() int {
	return s.mem[s.cur].min
}

func (s *Stack) pop() int {
	if !s.empty() {
		val := s.mem[s.cur].val
		s.cur--
		return val
	} else {
		panic("empty stack")
	}
}

func main() {
	stack := Stack{}
	stack.alloc(100)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println("min", stack.min())
	stack.push(-100)
	fmt.Println("min", stack.min())
	stack.pop()
	fmt.Println("min", stack.min())
	stack.pop()
}
