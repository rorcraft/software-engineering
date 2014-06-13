/*
 * In the classic problem of the Towers of Hanoi, you have 3 rods and N disks of different sizes
 * which can slide onto any tower. The puzzle starts with disks sorted in ascending order of
 * size from top to bottom (e.g., each disk sits on top of an even larger one).
 * You have the following constraints:
 *
 * - Only one disk can be moved at a time.
 * - A disk is slid off the top of one rod onto the next rod.
 * - A disk can only be placed on top of a larger disk.
 *
 * Write a program to move the disks from the first rod to the last using Stacks
 */
package main

import (
	"fmt"
)

type Stack struct {
	cur int
	mem []int
}

func (s *Stack) push(disk int) {
	if s.empty() || s.mem[s.cur] > disk {
		s.cur++
		s.mem[s.cur] = disk
	} else {
		panic("Cannot push larger disk on to the stack")
	}
}

func (s *Stack) empty() bool {
	return s.cur == -1
}

func (s *Stack) pop() int {
	if s.cur >= 0 {
		disk := s.mem[s.cur]
		s.mem[s.cur] = 0
		s.cur--
		return disk
	} else {
		panic("Empty stack")
	}
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.cur = -1
	s.mem = make([]int, size)
	return s
}

func NewHanoiTowser(size int) (*Stack, *Stack, *Stack) {
	s1 := NewStack(size)
	s2 := NewStack(size)
	s3 := NewStack(size)
	for i := size; i > 0; i-- {
		s1.push(i)
	}
	return s1, s2, s3
}

func PrintTower(s1, s2, s3 *Stack) {
	for i := range s1.mem {
		fmt.Println(s1.mem[i], s2.mem[i], s3.mem[i])
	}
	fmt.Println("--------------------------------")
}

func move(from, to) []*Stack {
}
// https://en.wikipedia.org/wiki/Tower_of_Hanoi#Solution
// Number the disks 1(ONE) through n (largest to smallest).
//
// If n is odd, the first move is from the Start to the Finish peg.
// If n is even, the first move is from the Start to the Using peg.
// Now, add these constraints:
//
// No odd disk may be placed directly on an odd disk.
// No even disk may be placed directly on an even disk.
// Never undo your previous move (that is, do not move a disk back to its immediate last peg).
// Considering those constraints after the first move, there is only one legal move at every subsequent turn.
func MoveHanoiTower(s1, s2, s3 *Stack) {
	n := len(s1.mem)
	lastMove := make([]*Stack, 2)

	if n % 2 == 1 {
		s3.push(s1.pop())

	} else {
		s2.push(s1.pop())
	}
	for !s1.empty() || !s2.empty() {

	}
	s2.push(s1.pop())
	PrintTower(s1, s2, s3)
	s3.push(s1.pop())
	s3.push(s2.pop())
	PrintTower(s1, s2, s3)
	s2.push(s1.pop())
	PrintTower(s1, s2, s3)
	s1.push(s3.pop())
	PrintTower(s1, s2, s3)
	s2.push(s3.pop())
	PrintTower(s1, s2, s3)
	s2.push(s1.pop())
	PrintTower(s1, s2, s3)
	s3.push(s1.pop())
	s3.push(s2.pop())
	s1.push(s2.pop())
	s1.push(s3.pop())
	s3.push(s2.pop())
	s2.push(s1.pop())
	s3.push(s1.pop())
	s3.push(s2.pop())
	PrintTower(s1, s2, s3)
}

func main() {
	s1, s2, s3 := NewHanoiTowser(6)
	MoveHanoiTower(s1, s2, s3)
}
