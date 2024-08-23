package sudokusolver

import (
	"testing"
)

var testSudoku [][]int
var solvedSudoku [][]int

func setUp() {
	testSudoku = [][]int{
		{4, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 2, 0, 8, 0, 5, 3, 0},
		{0, 0, 0, 7, 5, 0, 0, 0, 9},
		{0, 0, 5, 8, 7, 0, 0, 6, 2},
		{6, 0, 3, 9, 2, 0, 0, 8, 0},
		{0, 9, 0, 0, 6, 5, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0}, 
		{0, 0, 6, 0, 0, 7, 2, 0, 0}, 
		{0, 0, 0, 0, 9, 1, 7, 0, 3},
	}

	solvedSudoku = [][]int{
		{4, 5, 9, 3, 1, 6, 8, 2, 7},
        {7, 6, 2, 4, 8, 9, 5, 3, 1},
        {3, 8, 1, 7, 5, 2, 6, 4, 9},
        {1, 4, 5, 8, 7, 3, 9, 6, 2},
        {6, 7, 3, 9, 2, 4, 1, 8, 5},
        {2, 9, 8, 1, 6, 5, 3, 7, 4},
        {5, 1, 7, 2, 3, 8, 4, 9, 6},
        {9, 3, 6, 5, 4, 7, 2, 1, 8},
        {8, 2, 4, 6, 9, 1, 7, 5, 3},
	}
}

// Test case: A value can be safely inserted into the Sudoku table
func TestIsSafe(t *testing.T) {
	setUp()

	isSafe := isSafe(testSudoku, 0, 1, 5);
	want := true

	if isSafe != want {
		t.Errorf("Got %t but wanted: %t", isSafe, want)
	}
}

// Test case: A value cannot be inserted safely into the Sudoku table
func TestIsSafeNotSafe(t *testing.T) {
	setUp()

	isSafe := isSafe(testSudoku, 0, 1, 4);
	want := false

	if isSafe != want {
		t.Errorf("Got %t but wanted: %t", isSafe, want)
	}
}

// Test case: Validating valid Sudoku
func TestIsValidSudoku(t *testing.T) {
	setUp()

	isValid := IsValidSudoku(testSudoku)
	want := true

	if isValid != want {
		t.Errorf("Got %t but wanted: %t", isValid, want)
	}
}

// Test case: Validating invalid Sudoku
func TestIsValidSudokuNotValid(t *testing.T) {
	setUp()
	testSudoku[2][0] = 4

	isValid := IsValidSudoku(testSudoku)
	want := false

	if isValid != want {
		t.Errorf("Got %t but wanted: %t", isValid, want)
	}
}

// Test case: Solve the Sudoku
func TestSolveSudoku(t *testing.T) {
	setUp()

	isSolved := SolveSudoku(testSudoku)
	want := true

	if isSolved != want {
		t.Errorf("Got %t but wanted: %t", isSolved, want)
	}
}

// Test case: Sudoku cannot be solved
func TestSolveSudokuNotSolvable(t *testing.T) {
	setUp()
	testSudoku[2][0] = 4

	isSolved := SolveSudoku(testSudoku)
	want := false

	if isSolved != want {
		t.Errorf("Got %t but wanted: %t", isSolved, want)
	}
}