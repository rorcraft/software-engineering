/*
 * Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
 * Therefore, in real life, we would likely start a new stack when the previous stack
 * exceeds some threshold. Implement a data structure SetOfStacks that mimics this.
 * SetOfStacks should be composed of several stacks, and should create a new stack once
 * the previous one exceeds capacity. SetOfStacks.push() and SetOfStacks.pop() should behave
 * identically to a single stack (that is, pop() should return the same values as it would
 * if there were just a single stack).
 *
 * FOLLOW UP
 *
 * Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.
 */

package main
import (
	"fmt"
)

type SetOfStacks struct {
	stack *Stack
	threshold int
	n int
}

type Stack struct {
	cur int
	mem []int
	previous *Stack
}

func NewStack(size int, previous *Stack) *Stack {
	s := new(Stack)
	s.cur = -1
	s.mem = make([]int, size)
	s.previous = previous
	return s
}

func NewSetOfStacks(threshold int) *SetOfStacks {
	ss := new(SetOfStacks)
	ss.stack = NewStack(threshold, nil)
	ss.threshold = threshold
	ss.n = 1
	return ss
}

func (s *Stack) push(num int) {
	s.cur++
	s.mem[s.cur] = num
}

func (s *Stack) pop() int {
	val := s.mem[s.cur]
	s.cur--
	return val
}

func (s *SetOfStacks) push(num int) {
	if (s.stack.cur < s.threshold - 1) {
		s.stack.push(num)
	} else {
		s.stack = NewStack(s.threshold, s.stack)
		s.n++
		s.stack.push(num)
	}
}

func (s *SetOfStacks) pop() int {
	if (s.stack != nil && s.stack.cur == 0) {
		val := s.stack.pop()
		s.stack = s.stack.previous
		s.n--
		return val
	} else {
		return s.stack.pop()
	}
}

func main() {
	ss := NewSetOfStacks(3)
	for i := 1; i < 10; i++ {
		ss.push(i)
	}
	fmt.Println("n", ss.n)
	for i := 1; i < 10; i++ {
		fmt.Println(ss.pop())
	}
}
