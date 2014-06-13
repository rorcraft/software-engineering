/*
 * Write a program to sort a stack in ascending order.
 * You should not make any assumptions about how the stack is implemented.
 * The following are the only functions that should be used to write this program:
 * push | pop | peek | isEmpty.
 */

package main
import (
	"fmt"
	"container/heap"
	"math/rand"
)
// from example: http://golang.org/pkg/container/heap/
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // max heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func (s *Stack) peek() int {
	return s.mem[s.cur]
}

func shuffle(slice []int) []int {
	for i := len(slice) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		slice[i], slice[r] = slice[r], slice[i]
	}
	return slice
}

func sort(s *Stack) *Stack {
	h := new(IntHeap)
	heap.Init(h)
	for !s.empty() {
		heap.Push(h, s.pop())
	}
	for h.Len() > 0 {
		s.push(heap.Pop(h).(int))
	}
	return s
}

// using stacks only
func sortNaive(s *Stack) *Stack {
	sortedStack := NewStack(len(s.mem))
	for !s.empty() {
		n := s.pop()
		// always push the larger value to the bottom of the sorted stack
		for !sortedStack.empty() && n > sortedStack.peek() {
			s.push(sortedStack.pop())
		}
		sortedStack.push(n)
	}
	return sortedStack
}

func main() {
	s := NewStack(10)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = shuffle(a)
	fmt.Println("Shuffled", a)
	for _, e := range a {
		s.push(e)
	}
	s = sort(s)
	buffer := []int(nil)
	for !s.empty() {
		buffer = append(buffer, s.pop())
	}
	fmt.Println("Sorted", buffer)
	s2 := NewStack(10)
	for _, e := range a {
		s2.push(e)
	}
	s2 = sortNaive(s2)
	buffer = buffer[:0]
	for !s2.empty() {
		buffer = append(buffer, s2.pop())
	}
	fmt.Println("SortedNaive", buffer)
}
