/*
 * Given a sorted (increasing order) array, write an algorithm to
 * create a binary tree with minimal height.
 */

package main

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func BuildBinaryTree(ar []int) *Node {
	var buildBinaryTree func(start, end int) *Node
	buildBinaryTree = func(start, end int) *Node {
		if end <= start {
			return nil
		}
		mid := start + ((end - start) / 2)
		return &Node{
			Value: ar[mid],
			left:  buildBinaryTree(start, mid),
			right: buildBinaryTree(mid+1, end)}
	}
	return buildBinaryTree(0, len(ar))
}

func ToArray(t *Node, ar []int) []int {
	if t == nil {
		return ar
	}
	result := append(ToArray(t.left, ar), t.Value)
	result = append(result, ToArray(t.right, ar)...)
	return result
}

func Print(head *Node) {
	fmt.Println(ToArray(head, []int{}))
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tree := BuildBinaryTree(sorted)
	Print(tree)
	sorted = []int{1, 2, 3, 5, 5, 6, 7, 8, 20, 30, 40, 100, 110, 500}
	tree = BuildBinaryTree(sorted)
	Print(tree)
}
