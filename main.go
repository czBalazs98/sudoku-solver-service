package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bczmarko/sudoku-solver/sudokusolver"
)

type sudokuDto struct {
	Sudoku [][]int `json:"sudoku"`
}

func main() {
	router := gin.Default()
	router.POST("/solve-sudoku", solveSudoku)

	router.Run("localhost:8080")
}

func solveSudoku(c *gin.Context) {
	var sudoku sudokuDto

	// Check if request body can be binded
	if err := c.BindJSON(&sudoku); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Request is not valid"})
		return
	}

	// Check if the sudoku is valid
	if !sudokusolver.IsValidSudoku(sudoku.Sudoku) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Sudoku is not valid"})
		return
	}

	// Check if the sudoku can be solved
	if !sudokusolver.SolveSudoku(sudoku.Sudoku) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Sudoku cannot be solved"})
		return
	}

	// Solved sudoku in response
	c.IndentedJSON(http.StatusOK, sudoku)	
}