/*
 * Design an algorithm and write code to find the first common ancestor of two nodes in a
 * binary tree. Avoid storing additional nodes in a data structure.
 * NOTE: This is not necessarily a binary search tree.
 */

package main

import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// without parent
func LowestCommonAncestor(head, a, b *Node) *Node {
	if head == nil || head == a || head == b {
		return head
	}
	left := LowestCommonAncestor(head.Left, a, b)
	right := LowestCommonAncestor(head.Right, a, b)

	if left != nil && right != nil {
		return head
	} else if left != nil {
		return left
	} else {
		return right
	}
}

// with parent

// dynamic programming storing depths

func main() {
	fmt.Println("4.6 Lowest Common Ancestor")
	/*

	           0
	          /  \
	         /    \
	        /      \_____
	       1.L    1.R   2.R
	      /  \    /      / \
	     2    C  2.L   3.L  3.R
	    /  \    /  \   /\   /\
	   A    B  D    E F  G H  J
	*/
	head := &Node{Value: "0"}
	head.Left = &Node{Value: "1 left"}
	head.Right = &Node{Value: "1 right"}
	Right1 := head.Right
	head.Right.Left = &Node{Value: "2 left"}
	head.Right.Right = &Node{Value: "2 right"}
	head.Right.Left.Left = &Node{Value: "D"} //D
	D := head.Right.Left.Left
	head.Right.Left.Right = &Node{Value: "E"} //E
	head.Right.Right.Left = &Node{Value: "3 left"}
	head.Right.Right.Left.Left = &Node{Value: "F"} //F
	F := head.Right.Right.Left.Left
	head.Right.Right.Left.Right = &Node{Value: "G"} //G
	head.Right.Right.Right = &Node{Value: "3 right"}
	head.Right.Right.Right.Left = &Node{Value: "H"}  //H
	head.Right.Right.Right.Right = &Node{Value: "J"} //J
	head.Left.Left = &Node{Value: "2 left"}
	head.Left.Right = &Node{Value: "C"}      //C
	head.Left.Left.Left = &Node{Value: "A"}  //A
	head.Left.Left.Right = &Node{Value: "B"} //B
	fmt.Println(D.Value, F.Value, "LCA:", LowestCommonAncestor(head, D, F).Value)
	fmt.Println(D.Value, D.Value, "LCA:", LowestCommonAncestor(head, D, D).Value)
	fmt.Println(D.Value, Right1.Value, "LCA:", LowestCommonAncestor(head, D, Right1).Value)
	fmt.Println(Right1.Value, D.Value, "LCA:", LowestCommonAncestor(head, Right1, D).Value)
}
