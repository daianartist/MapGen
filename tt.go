package main

import (
	"fmt"
	"github.com/alem-platform/ap"
)

func putN(r rune, n int) {
	for i := 0; i < n; i++ {
		ap.PutRune(r)
	}
}

func putCentered(sym rune) {
	putN(' ', 3)
	ap.PutRune(sym)
	putN(' ', 3)
}

func drawTopBorder(cols int) {
	ap.PutRune(' ')
	putN('_', cols*7+cols-1)
}

func drawRowBody(row []int) {
	for line := 0; line < 2; line++ {
		ap.PutRune('\n')
		ap.PutRune('|')
		for i := 0; i < len(row); i++ {
			switch row[i] {
			case 0:
				putN('X', 7)
			case 1:
				putN(' ', 7)
			case 2:
				if line == 1 {
					putCentered('>')
				} else {
					putN(' ', 7)
				}
			case 3:
				if line == 1 {
					putCentered('@')
				} else {
					putN(' ', 7)
				}
			}
			ap.PutRune('|')
		}
	}
}

func drawBottomRow(row []int) {
	ap.PutRune('\n')
	ap.PutRune('|')
	for i := 0; i < len(row); i++ {
		if row[i] == 0 {
			putN('X', 7)
		} else {
			putN('_', 7)
		}
		ap.PutRune('|')
	}
}

func printError(msg string) {
	for _, r := range msg {
		ap.PutRune(r)
	}
	ap.PutRune('\n')
}

func main() {
	var rows, cols int
	var count int
	var err error

	count, err = fmt.Scanf("%d %d\n", &rows, &cols)
	if err != nil || count != 2 {
		printError("Error: Invalid input. Please provide two numbers for rows and cols.")
		return
	}

	if rows <= 0 || cols <= 0 {
		printError("Error: Grid dimensions must be positive numbers.")
		return
	}

	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		var line string
		count, err = fmt.Scanf("%s\n", &line)
		if err != nil || count != 1 {
			printError("Error: Could not read grid row.")
			return
		}

		if len(line) != cols {
			printError("Error: Row has unexpected length")
			return
		}

		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			char := line[j]
			if char == '0' {
				grid[i][j] = 0
			} else if char == '1' {
				grid[i][j] = 1
			} else if char == '2' {
				grid[i][j] = 2
			} else if char == '3' {
				grid[i][j] = 3
			} else {
				printError("Error: Invalid character in grid!")
				return
			}
		}
	}

	drawTopBorder(cols)
	for i := 0; i < rows; i++ {
		drawRowBody(grid[i])
		drawBottomRow(grid[i])
	}
	ap.PutRune('\n')
}