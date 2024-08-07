package sudokusolver

import "fmt"

// Check if a number can be inserted to the given position in a sudoku table
func isSafe(sudoku [][]int, row int, col int, num int) bool {
	
	// Checking row for match
	for i := 0; i < len(sudoku); i++ {
		if sudoku[row][i] == num {
			return false
		}
	}

	// Checking column for match
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i][col] == num {
			return false
		}
	}

	// Checking the box for match
	boxRowStart := row - row % 3
	boxColStart := col - col % 3
	for i := boxRowStart; i < boxRowStart + 3; i++ {
		for j := boxColStart; j < boxColStart + 3; j++ {
			if sudoku[i][j] == num {
				return false
			}
		}
	}

	// No match found, number can be inserted
	return true
}

// Solving the given sudoku
func SolveSudoku(sudoku [][]int) bool {
	row, col := -1, -1
	isEmpty := false

	// Checking for empty field
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku); j++ {
			if sudoku[i][j] == 0 {
				row, col = i, j
				isEmpty = true
				break;
			}
		}

		if isEmpty {
			break
		}
	}

	// If no more empty fields, the sudoku is solved
	if !isEmpty {
		return true
	}

	// Backtracking
	for num := 1; num <= 9; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num
			if SolveSudoku(sudoku) {
				return true
			} else {
				sudoku[row][col] = 0
			}
		}
	}
	
	// If sudoku cannot be solved return false
	return false
}

// Print the sudoku table
func PrintSudoku(sudoku [][]int) {
	for i := range sudoku {
		for j := range sudoku[i] {
			fmt.Printf("%d ", sudoku[i][j])
		}
		fmt.Println()
	}
}