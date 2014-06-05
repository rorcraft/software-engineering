package main
/* Question
 * write a method to replace all spaces in a string with ‘%20’.
 */

import (
	"bytes"
	"fmt"
)

func replace(str string) string {
	var buffer bytes.Buffer
	for _, c := range str {
		if string(c) == " " {
			buffer.WriteString("%20")
		} else {
			buffer.WriteString(string(c))
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println("hello world", "replaced", replace("hello world"))
}
