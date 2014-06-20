/*
 * Given a binary search tree, design an algorithm which creates a linked list of all
 * the nodes at each depth (i.e., if you have a tree with depth D, you’ll have D linked lists).
 */

package main
import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func BuildList(n *Node) [][]string {
	lists := [][]string{}
	var buildList func (n *Node, h int)
	buildList = func (n *Node, h int) {
		if n == nil {
			return
		}
		if len(lists) <= h {
			lists = append(lists, []string{})
		}
		lists[h] = append(lists[h], n.Value)
		buildList(n.Left, h+1)
		buildList(n.Right, h+1)
	}
	buildList(n, 0)
	return lists
}

func Print(lists [][]string) {
	for i, level := range lists {
		fmt.Println(i, level)
	}
}
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
func main() {
	head := &Node{Value: "0"}
	head.Left = &Node{Value: "1 left"}
	head.Right = &Node{Value: "1 right"}
	head.Right.Left = &Node{Value: "2 left"}
	head.Right.Right = &Node{Value: "2 right"}
	head.Right.Left.Left = &Node{Value: "D"} //D
	head.Right.Left.Right = &Node{Value: "E"} //E
	head.Right.Right.Left = &Node{Value: "3 left"}
	head.Right.Right.Left.Left = &Node{Value: "F"} //F
	head.Right.Right.Left.Right = &Node{Value: "G"} //G
	head.Right.Right.Right = &Node{Value: "3 right"}
	head.Right.Right.Right.Left = &Node{Value: "H"} //H
	head.Right.Right.Right.Right = &Node{Value: "J"} //J
	head.Left.Left = &Node{Value: "2 left"}
	head.Left.Right = &Node{Value: "C"} //C
	head.Left.Left.Left = &Node{Value: "A"} //A
	head.Left.Left.Right = &Node{Value: "B"} //B
	fmt.Println("###########################################")
	fmt.Println(" 4.4 linked list per depth of binary tree")
	fmt.Println("###########################################")
	lists := BuildList(head)
	Print(lists)
}
