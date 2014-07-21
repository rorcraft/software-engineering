// Implement the “paint fill” function that one might see on many image editing programs.
// That is, given a screen (represented by a 2-dimensional array of Colors), a point,
// and a new color, fill in the surrounding area until you hit a border of that color.
package main

import "fmt"

// 1. use bfs

// 2. recursive
type Color int
const (
	White Color = iota
	Black
	Red
	Green
	Blue
)

func paintFill(screen [][]Color, x, y int, c Color, initColor Color) {
	if initColor < 0 {
		initColor = screen[x][y]
	}

	if x < 0 || x >= len(screen[0]) || y < 0 || y >= len(screen) {
		return
	} else if screen[x][y] == initColor {
		screen[x][y] = c
		paintFill(screen, x-1, y, c, initColor)
		paintFill(screen, x+1, y, c, initColor)
		paintFill(screen, x, y-1, c, initColor)
		paintFill(screen, x, y+1, c, initColor)
	}
}

func setupScreen(x, y int) [][]Color {
	screen := make([][]Color, x)
	for i := x - 1; i >= 0; i-- {
		screen[i] = make([]Color, y)
	}

	// paint the border
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || i == x-1 || j == 0 || j == y-1 {
				screen[i][j] = Green
			}
		}
	}
	return screen
}

func printScreen(screen [][]Color) {
	x := len(screen)
	y := len(screen[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	screen := setupScreen(5, 5)
	fmt.Println("8.6 paintfill")
	printScreen(screen)
	paintFill(screen, 2, 3, Red, -1)
	printScreen(screen)
}
