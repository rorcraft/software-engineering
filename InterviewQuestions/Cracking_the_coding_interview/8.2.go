// Imagine a robot sitting on the upper left hand corner of an NxN grid.
// The robot can only move in two directions: right and down.
// How many possible paths are there for the robot?
//
// FOLLOW UP
//
// Imagine certain squares are “off limits”, such that the robot can not step on them.
// Design an algorithm to get all possible paths for the robot.
package main

import "fmt"

// Mathematical Solution: Take m-1 combinations from m-1+n-1 steps.
// C(m-1+n-1, m-1)=(m-1+n-1)! / ( (m-1)! * (n-1)! )


// similar to tower of hanoi problem, work backwards recursively.
func path(n, m int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return path(n-1, m) + path(n, m-1)
}

func pathWithLimits(limits [][]bool, n, m int) int {
	if limits[n-1][m-1] == false {
		return 0
	} else if n == 1 || m == 1 {
		return 1
	}
	return pathWithLimits(limits, n-1, m) + pathWithLimits(limits, n, m-1)
}

func makeMap() [][]bool {
  m := make([][]bool, 5)
	m[0] = []bool{true,true,true,false,true}
	m[1] = []bool{true,true,true,true,true}
	m[2] = []bool{true,true,true,true,true}
	m[3] = []bool{true,true,true,true,true}
	m[4] = []bool{true,false,true,true,true}
	return m
}

func main() {
	fmt.Println("5,5", path(5,5))
	fmt.Println("5,5 with limits", pathWithLimits(makeMap(), 5,5))
}
