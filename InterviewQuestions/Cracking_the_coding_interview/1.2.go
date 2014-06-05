package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
 * Question
 * Write code to reverse a C-Style String.
 * (C-String means that “abcd” is represented as five characters, including the null character.)
 */

/*
 * Note:
 * C-style string has \0 char for termination
 */

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

// this does not work as golang strings are immutable
/*
func failed_reverse(str string) string {
	n := len(str)
	end := n / 2
	for i := 0; i < end; i++ {
		str[i], str[n-i] = str[n-i], str[i] // cannot assign str[i], str[n-i]
	}
	return str
}
*/

func reverse(str string) string {
	out := make([]rune, utf8.RuneCountInString(str))
	n := len(str)
	for _, c := range str {
		n--
		out[n] = c
	}
	return string(out)
}

// bonus: http://www.techinterview.org/post/526374214/reverse-a-string
/*
 * reverse words. e.g. "hello world" => "world hello"
 */

func reverseWord(str string) string {
	reversed := strings.Split(reverse(str), " ")
	for i, word := range reversed {
		reversed[i] = reverse(word)
	}
	return strings.Join(reversed, " ")
}

func main() {
	str := "hello world"
	fmt.Println("reverse", str, "=>", reverse(str))
	fmt.Println("reverse_bonus", str, "=", reverseWord(str))
}
