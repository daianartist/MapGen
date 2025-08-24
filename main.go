package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func drawTopBorder(cols int) {
	ap.PutRune(' ')
	totalUnderscores := cols * 7
	for i := 0; i < totalUnderscores+cols-1; i++ {
		ap.PutRune('_')
	}
}

func drawRowBody(row []int) {
	for line := 0; line < 2; line++ {
		ap.PutRune('\n')
		ap.PutRune('|')
		for i := 0; i < len(row); i++ {
			cell := row[i]
			switch cell {
			case 0:
				//  X
				for k := 0; k < 7; k++ {
					ap.PutRune('X')
				}
			case 1:
				// nothing
				for k := 0; k < 7; k++ {
					ap.PutRune(' ')
				}
			case 2:
				//  >
				if line == 1 {
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune('>')
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune(' ')
				} else {
					for k := 0; k < 7; k++ {
						ap.PutRune(' ')
					}
				}
			case 3:
				//  @
				if line == 1 {
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune('@')
					ap.PutRune(' ')
					ap.PutRune(' ')
					ap.PutRune(' ')
				} else {
					for k := 0; k < 7; k++ {
						ap.PutRune(' ')
					}
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
		cell := row[i]
		switch cell {
		case 0:
			for k := 0; k < 7; k++ {
				ap.PutRune('X')
			}
		default:
			for k := 0; k < 7; k++ {
				ap.PutRune('_')
			}
		}
		ap.PutRune('|')
	}
}
func errManager(err_text string){
	runes := []rune(err_text) 
	for i := 0; i < len(runes); i++ {
		ap.PutRune(runes[i])
	}
}

func main() {
	var rows, cols int
	if _, err := fmt.Scanf("%d %d\n", &rows, &cols); err != nil {
		errManager("Error: Invalid input. Please provide two numbers for rows and cols.")
		return //graceful exit
	}
	if rows <= 0 || cols <= 0 {
		errManager("Error: Grid dimensions must be positive numbers.")
		return//graceful exit
	}
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		var line string
		
		if n, err := fmt.Scanf("%s\n", &line); err != nil || n == 0 {
			errManager("Error: Could not read grid row.")
			return
		}
		if len(line) != cols {
			errManager("Error: Row has unexpected length") 
			return//graceful exit
		}
		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			switch line[j] {
			case '0':
				grid[i][j] = 0
			case '1':
				grid[i][j] = 1
			case '2':
				grid[i][j] = 2
			case '3':
				grid[i][j] = 3
			default:
				errManager("Error: Invalid character in grid!")
				return//graceful exit
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