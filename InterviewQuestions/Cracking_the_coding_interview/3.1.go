/*
 * Describe how you could use a single array to implement three stacks.
 */

// http://hawstein.com/posts/3.1.html

// this version just divides the entire array evenly and create 3 continuous stacks.
package main
import (
	"fmt"
)

type Stack struct {
	start int
	index int
	maxSize int
	mem []string
}

func (s *Stack) push(content string) {
	if s.index - s.start < s.maxSize {
		s.mem[s.index] = content
		s.index++
	} else {
		panic("stack overflow")
	}
}

func (s *Stack) pop() string {
	if s.index - s.start > 0 {
		s.index--
		str := s.mem[s.index]
		s.mem[s.index] = ""
		return str
	} else {
		return ""
	}
}

func allocate(size int) (*Stack, *Stack, *Stack) {
	fmt.Println("allocating ", size)
	memory := make([]string, size)
	stackSize := size / 3
	stacks := make([]*Stack, 3)
	for i := 0; i < 3; i++ {
		index := stackSize * i
		stacks[i] = &Stack{start: index, index: index, maxSize: stackSize, mem: memory}
	}
	return stacks[0], stacks[1], stacks[2]
}

func main() {
	stack1, stack2, stack3 := allocate(9)
	stack1.push("hello")
	stack1.push("world")
	stack2.push("hello")
	stack2.push("a")
	stack2.push("b")
	// stack2.push("c")
	stack3.push("hello")
	stack3.push("i was here")
	fmt.Println("popped", stack3.pop())
	fmt.Println("popped", stack3.pop())

	for i, s := range stack1.mem {
		fmt.Println(i, s)
	}
}
