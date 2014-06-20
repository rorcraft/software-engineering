package digraph

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_ParseNewDigraph(test *testing.T) {
	file, _ := os.Open("digraph1.txt")
	reader := bufio.NewReader(file)
	dig := ParseNewDigraph(reader)
	fmt.Println(dig.String())
}
