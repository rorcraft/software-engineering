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
	size int
}

func (s *Stack) push(disk int) {
	if s.empty() || s.mem[s.cur] > disk {
		s.cur++
		s.size++
		s.mem[s.cur] = disk
	} else {
		panic("Cannot push larger disk on to the stack")
	}
}

func (s *Stack) empty() bool {
	return s.cur == -1
}

func (s *Stack) peek() int {
	return s.mem[s.cur]
}

func (s *Stack) pop() int {
	if s.cur >= 0 {
		disk := s.mem[s.cur]
		s.mem[s.cur] = 0
		s.cur--
		s.size--
		return disk
	} else {
		panic("Empty stack")
	}
}

func NewStack(size int) *Stack {
	s := new(Stack)
	s.cur = -1
	s.mem = make([]int, size)
	s.size = 0
	return s
}

type HanoiTower struct {
	s1 *Stack
	s2 *Stack
	s3 *Stack
	N  int
}

func NewHanoiTower(size int) *HanoiTower {
	ht := new(HanoiTower)
	ht.s1 = NewStack(size)
	ht.s2 = NewStack(size)
	ht.s3 = NewStack(size)
	for i := size; i > 0; i-- {
		ht.s1.push(i)
	}
	ht.N = size
	return ht
}

func (ht *HanoiTower) Print() {
	for i := 0 ; i < ht.N; i++ {
		fmt.Println(ht.s1.mem[i], ht.s2.mem[i], ht.s3.mem[i])
	}
	fmt.Println("--------------------------------")
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

// this never finished.
func (ht *HanoiTower) Solve() {
	lastMove := make([]*Stack, 2)

	move := func(from, to *Stack) {
		to.push(from.pop())
		lastMove[0] = from
		lastMove[1] = to
	}

	validMove := func(from, to *Stack) bool {
		isUndo := lastMove[0] == to && lastMove[1] == from
		return !from.empty() &&
			(to.empty() ||
			(!to.empty() && (from.peek()%2 != to.peek()%2) && (from.peek() < to.peek()))) &&
			!isUndo
	}

	possibleMoves := make([][]*Stack, 6)
	possibleMoves[0] = []*Stack{ht.s1, ht.s2}
	possibleMoves[1] = []*Stack{ht.s1, ht.s3}
	possibleMoves[2] = []*Stack{ht.s2, ht.s3}
	possibleMoves[3] = []*Stack{ht.s2, ht.s1}
	possibleMoves[4] = []*Stack{ht.s3, ht.s1}
	possibleMoves[5] = []*Stack{ht.s3, ht.s2}

	if ht.s1.peek()%2 == 1 {
		move(ht.s1, ht.s3)
	} else {
		move(ht.s1, ht.s2)
	}
	for ht.s3.size <= ht.N {
		for _, possibleMove := range possibleMoves {
			if validMove(possibleMove[0], possibleMove[1]) {
				move(possibleMove[0], possibleMove[1])
				ht.Print()
			}
		}
	}
}

// Recursive solution
// A key to solving this puzzle is to recognize that it can be solved by breaking the problem down into a collection of smaller problems and further breaking those problems down into even smaller problems until a solution is reached. For example:
//
// label the pegs A, B, C — these labels may move at different steps
// let n be the total number of discs
// number the discs from 1 (smallest, topmost) to n (largest, bottommost)
// To move n discs from peg A to peg C:
//
// move n−1 discs from A to B. This leaves disc n alone on peg A
// move disc n from A to C
// move n−1 discs from B to C so they sit on disc n
func (ht *HanoiTower) SolveRecur() {
	move := func(from, to *Stack) {
		to.push(from.pop())
		ht.Print()
	}
	var recur func (n int, src, bri, dst *Stack)
	recur = func (n int, src, bri, dst *Stack) {
		if (n == 1) {
			move(src, dst)
		} else {
			recur(n-1, src, dst, bri)
			move(src, dst)
			recur(n-1, bri, src, dst)
		}
	}
	recur(ht.N, ht.s1, ht.s2, ht.s3)
}

func main() {
	ht := NewHanoiTower(6)
	// ht.Solve()
	ht.SolveRecur()
}
