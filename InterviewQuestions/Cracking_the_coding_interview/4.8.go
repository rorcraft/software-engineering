/*
 * You are given a binary tree in which each node contains a value.
 * Design an algorithm to print all paths which sum up to that value.
 * Note that it can be any path in the tree - it does not have to start at the root.
 */

package main
import (
	"strings"
	"fmt"
	"strconv"
)


type Node struct {
	Value int
	Left *Node
	Right *Node
	Parent *Node
}

func FindSum(head *Node, sum int) []*Node {
	paths := []*Node{}
	var findSum func (head *Node, remainderSum int)
	findSum = func (head *Node, remainderSum int) {
		if head == nil || remainderSum < 0 {
			return
		}
		remainderSum = remainderSum - head.Value
		if (remainderSum == 0) {
			paths = append(paths, head)
		}
		findSum(head.Left, remainderSum)
		findSum(head.Right, remainderSum)
	}
	findSum(head, sum)
	return paths
}

func PrintPaths(paths []*Node) {
	for _, leaf := range paths {
		path := []int{}
		for n := leaf; n != nil; n = n.Parent {
			path = append(path, n.Value)
		}
		// reverse
		reversed := make([]string, len(path))
		size := len(path)
		for i := 0; i < size; i++ {
			reversed[i] = strconv.Itoa(path[size - i - 1])
		}
		fmt.Println(strings.Join(reversed, " -> "))
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
	fmt.Println("4.8 find path for sum")
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tree := BuildBinaryTree(sorted)
	paths := FindSum(tree, 12)
	PrintPaths(paths)
}
