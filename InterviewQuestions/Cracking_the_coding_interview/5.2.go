/*
 * Given a (decimal - e.g. 3.72) number that is passed in as a string,
 * print the binary representation. If the number can not be represented accurately in binary,
 * print “ERROR”.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FloatToBin(in string) string {
	parts := strings.Split(in, ".")
	intPart, _ := strconv.ParseInt(parts[0], 10, 64)
	var fracPart float64
	if (len(parts) > 0) {
		fracPart, _ = strconv.ParseFloat("0."+parts[1], 64)
	}
	intBinStr := fmt.Sprintf("%b", intPart)
	fracBinStr := ""
	for fracPart > 0 {
		if len(fracBinStr) > 64 {
			// fmt.Println("ERROR")
			return "ERROR"
		}
		fracPart *= 2
		if fracPart >= 1 {
			fracBinStr += "1"
			fracPart -= 1
		} else {
			fracBinStr += "0"
		}
	}
	return fmt.Sprintf("%s.%s", intBinStr, fracBinStr)
}
// not the same as the question http://kipirvine.com/asm/workbook/floating_tut.htm
func main() {
	fmt.Println("5.2 - float to binary")
	str := "1.6875"
	fmt.Printf("%s -> %s\n", str, FloatToBin(str))
	str = "1.25"
	fmt.Printf("%s -> %s\n", str, FloatToBin(str))
}
