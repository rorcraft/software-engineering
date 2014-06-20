/*
 * Implement a function to check if a tree is balanced.
 * For the purposes of this question, a balanced tree is defined to be a tree such that
 * no two leaf nodes differ in distance from the root by more than one.
 */

// https://stackoverflow.com/questions/742844/how-to-determine-if-binary-tree-is-balanced
/*

not balanced, C and F has diff in height 2 (>1)
        /\
       /  \
      /    \
     /      \_____
    /\      /     \_
   /  \    /      / \
  /\   C  /\     /   \
 /  \    /  \   /\   /\
A    B  D    E F  G H  J

> A tree where the maximum height of any branch is no more
  than one more than the minimum height of any branch.
*/
package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced(tree *Node) bool {
	if (tree == nil) {
		return false
	}
	ch := make(chan int)
	doneCheck := make(chan bool, 1)
	continueCheck := make(chan bool, 1)
	finished := false
	defer close(doneCheck)
	defer close(continueCheck)

	var getHeight func(t *Node, h int)
	getHeight = func(t *Node, h int) {
		if finished {
			return
		}
		if t.Left == nil && t.Right == nil {
			fmt.Println("found leaf", h)
			ch<-h
			finished = !<-continueCheck
			return
		}
		if t.Left != nil {
			getHeight(t.Left, h+1)
		}

		if t.Right != nil {
			getHeight(t.Right, h+1)
		}
	}

	var checkHeight = func(heights chan int) {
		min, max := -1 , -1
		balanced := true
		for {
			h, ok := <-heights
			if (!ok) { //exhausted
				doneCheck<-balanced
				break
			}
			// fmt.Println("got height", h)
			if min < 0 && max < 0 {
				min, max = h, h
			} else if h < min {
				min = h
			} else if h > max {
				max = h
			}
			if Abs(max-min) > 1 {
				balanced = false
				continueCheck<-false
			} else {
				continueCheck<-true
			}
		}
	}
	go checkHeight(ch)
	go func () {
		getHeight(tree, 0)
		close(ch)
	}()
	return <-doneCheck
}

func main() {
	head := &Node{}
	head.Left = &Node{}
	head.Right = &Node{}
	fmt.Println("isBalanced?", isBalanced(head))
//         /\
//        /  \
//       /    \
//      /      \_____
//     /\      /     \_
//    /  \    /      / \
//   /\   C  /\     /   \
//  /  \    /  \   /\   /\
// A    B  D    E F  G H  J
	head.Right.Left = &Node{}
	head.Right.Right = &Node{}
	head.Right.Left.Left = &Node{} //D
	head.Right.Left.Right = &Node{} //E
	head.Right.Right.Left = &Node{}
	head.Right.Right.Left.Left = &Node{} //F
	head.Right.Right.Left.Right = &Node{} //G
	head.Right.Right.Right = &Node{}
	head.Right.Right.Right.Left = &Node{} //H
	head.Right.Right.Right.Right = &Node{} //J
	head.Left.Left = &Node{}
	head.Left.Right = &Node{} //C
	head.Left.Left.Left = &Node{} //A
	head.Left.Left.Right = &Node{} //B
	fmt.Println("isBalanced?", isBalanced(head))
	head.Left.Right.Left = &Node{} //C+left
	fmt.Println("isBalanced?", isBalanced(head))
}
