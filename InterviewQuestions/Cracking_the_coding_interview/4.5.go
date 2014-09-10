/*
 * Write an algorithm to find the ‘next’ node (i.e., in-order successor) of a given node
 * in a binary search tree where each node has a link to its parent.
 */
package main

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func FindNext(node *Node) *Node {
	if node.Right == nil {
		// need to recur up to find the next large
		// if node is right of parent, parent is smaller.
		return node.Parent
	} else {
		next := node.Right
		for next.Left != nil {
			next = next.Left
		}
		return next
	}
}

// modified from 4.3
func BuildBinaryTree(ar []int) *Node {
	var buildBinaryTree func(start, end int, parent *Node) *Node
	buildBinaryTree = func(start, end int, parent *Node) *Node {
		if end <= start {
			return nil
		}
		mid := start + ((end - start) / 2)
		this := &Node{Value: ar[mid], Parent: parent}
		this.Left = buildBinaryTree(start, mid, this)
		this.Right = buildBinaryTree(mid+1, end, this)
		return this
	}
	return buildBinaryTree(0, len(ar), nil)
}

func main() {
	fmt.Println("4.5 find 'next' node")
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tree := BuildBinaryTree(sorted)
	next := FindNext(tree)
	fmt.Println(tree.Value, "next is:", next.Value)
	node := tree.Left.Left
	next = FindNext(node)
	fmt.Println(node.Value, "next is:", next.Value)
	node = tree.Right.Left
	next = FindNext(node)
	fmt.Println(node.Value, "next is:", next.Value)
}
