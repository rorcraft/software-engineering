/*
 * Given a directed graph, design an algorithm to find out whether there is a route
 * between two nodes.
 */

package main
import (
	"fmt"
	"os"
	"bufio"
	"./lib/digraph"
)

func route(dig *digraph.Digraph, src, dst int) bool {
	queue := make([]int, 1)
	visited := make([]bool, dig.V)
	queue[0] = src
	for len(queue) > 0 {
		v := queue[0]
		if len(queue) > 1 {
			queue = queue[1:len(queue)-1]
		} else {
			queue = queue[:0]
		}
		if v == dst {
			return true
		} else {
			visited[v] = true
			for _, adj := range dig.Adj[v] {
				if !visited[adj] {
					queue = append(queue, adj)
				}
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("./lib/digraph/digraph1.txt")
	reader := bufio.NewReader(file)
	dig := digraph.ParseNewDigraph(reader)
	// fmt.Println(dig.String())
	fmt.Println("route?", route(dig, 1, 4))
	fmt.Println("route?", route(dig, 4, 1)) // directed
	fmt.Println("route?", route(dig, 7, 0))
	fmt.Println("route?", route(dig, 7, 6))
/*
 $ go run 4.2.go
route? false
route? true
route? true
route? false
*/
}
