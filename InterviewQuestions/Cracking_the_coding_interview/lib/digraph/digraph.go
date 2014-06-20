// http://algs4.cs.princeton.edu/42directed/Digraph.java.html
package digraph
import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

type Digraph struct {
	V   int
	E   int
	Adj [][]int
}

func NewDigraph(V int) *Digraph {
	dig := new(Digraph)
	dig.V = V
	dig.E = 0
	dig.Adj = make([][]int, V)
	for i, _ := range dig.Adj {
		dig.Adj[i] = []int{}
	}
	return dig
}

func ParseNewDigraph(r *bufio.Reader) *Digraph {
	readInt := func(delim byte) int {
		var str string
		if delim == '\n' {
			str, _ = r.ReadString(delim)
		} else {
			for {
				str, _ = r.ReadString(delim)
				if str != " " {
					break
				}
			}
		}
		re := regexp.MustCompile(`\d+`)
		stri := re.FindString(str)
		i, _ := strconv.Atoi(stri)
		return i
	}
	dig := new(Digraph)
	dig.V = readInt('\n')
	if dig.V < 0 {
		panic("Number of vertices in a Digraph must be nonnegative")
	}
	dig.Adj = make([][]int, dig.V)
	for i, _ := range dig.Adj {
		dig.Adj[i] = []int{}
	}
	E := readInt('\n')
	if E < 0 {
		panic("Number of edges in a Digraph must be nonnegative")
	}

	for i := 0; i < E; i++ {
		dig.AddEdge(readInt(' '), readInt('\n'))
	}
	return dig
}

func (dig *Digraph) AddEdge(v, w int) {
	if v < 0 || v >= dig.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, (dig.V - 1)))
	}
	if w < 0 || w >= dig.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", w, (dig.V - 1)))
	}
	dig.Adj[v] = append(dig.Adj[v], w)
	dig.E++
}

/**
 * Returns a string representation of the graph.
 */
func (dig *Digraph) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d vertices %d edges\n", dig.V, dig.E))
	for i, v := range dig.Adj {
		buffer.WriteString(fmt.Sprintf("%d: ", i))
		for _, w := range v {
			buffer.WriteString(fmt.Sprintf("%d ", w))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
