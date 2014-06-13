/*
 * Implement a MyQueue class which implements a queue using two stacks.
 */

package main
import (
	"fmt"
	"strconv"
)

type MyQueue struct {
	enqueueStack *Stack
	dequeueStack *Stack
	enqueueMode bool
}

type Stack struct {
	cur int
	mem []int
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.cur = -1
	s.mem = make([]int, size)
	return s
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

func (s *Stack) empty() bool {
	return s.cur == -1
}

func (q *MyQueue) enqueue(data int) {
	if (!q.enqueueMode) {
		for !q.dequeueStack.empty() {
			q.enqueueStack.push(q.dequeueStack.pop())
		}
		q.enqueueMode = true
	}
	q.enqueueStack.push(data)
}

func (q *MyQueue) dequeue(data int) int {
	if (q.enqueueMode) {
		for !q.enqueueStack.empty() {
			q.dequeueStack.push(q.enqueueStack.pop())
		}
		q.enqueueMode = false
	}
	return q.dequeueStack.pop()
}

func NewQueue() *MyQueue {
	q := new(MyQueue)
	q.enqueueStack = NewStack(10)
	q.dequeueStack = NewStack(10)
	q.enqueueMode = true
	return q
}

func main() {
	q := NewQueue()
	buffer := []string(nil)
	fmt.Println("MyQueue should be FIFO")
	for i := 0; i < 10; i++ {
		q.enqueue(i)
		buffer = append(buffer, strconv.Itoa(i))
	}
	fmt.Println("Enqueued", buffer)
	buffer = buffer[:0]
	for i := 0; i < 10; i++ {
		buffer = append(buffer, strconv.Itoa(q.dequeue(i)))
	}
	fmt.Println("Dequeued", buffer)
}
