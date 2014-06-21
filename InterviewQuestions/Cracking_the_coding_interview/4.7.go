/*
 * You have two very large binary trees: T1, with millions of nodes, and T2,
 * with hundreds of nodes. Create an algorithm to decide if T2 is a subtree of T1
 */

package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func IsSubtree(a, b *Node) bool {
	matched := FindMatch(a, b)
	if matched == nil {
		return false
	}
	return IsIdentifical(matched, b)
}

func FindMatch(a, b *Node) *Node {
	if a == nil || a.Value == b.Value {
		return a
	}
	left := FindMatch(a.Left, b)
	right := FindMatch(a.Right, b)
	if left != nil {
		return left
	} else {
		return right
	}
}

func IsIdentifical(a, b *Node) bool {
	if a == nil || b == nil || a.Value != b.Value {
		return false
	}
	identical := true
	if a.Left != nil || b.Left != nil {
		identical = identical && IsIdentifical(a.Left, b.Left)
	}
	if a.Right != nil || b.Right != nil {
		identical = identical && IsIdentifical(a.Right, b.Right)
	}
	return identical
}

/************************************
       Tree S
          10
        /    \
      4       6
       \
        30


        Tree T
              26
            /   \
          10     3
        /    \     \
      4       6      3
       \
        30
**************************************/
func main() {
	fmt.Println("4.7 Check subtree of a binary tree")
	T := &Node{Value: 26}
	T.Left = &Node{Value: 10}
	T.Right = &Node{Value: 3}
	T.Left.Left = &Node{Value: 4}
	T.Left.Left.Right = &Node{Value: 30}
	T.Left.Right = &Node{Value: 6}
	T.Right.Right = &Node{Value: 3}

	S := &Node{Value: 10}
	S.Left = &Node{Value: 4}
	S.Left.Right = &Node{Value: 30}
	S.Right = &Node{Value: 6}

	fmt.Println("Is S a subtree of T?", IsSubtree(T, S)) //true
	S.Left.Left = &Node{Value: 4} // extra node
	fmt.Println("Is S a subtree of T?", IsSubtree(T, S)) //false

	S = &Node{Value: 30} // single leaf node
	fmt.Println("Is S a subtree of T?", IsSubtree(T, S)) //true
}
