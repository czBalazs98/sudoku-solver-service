package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var validSudokuMock = [][]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

var solvedSudokuMock = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2}, 
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7}, 
	{8, 5, 9, 7, 6, 1, 4, 2, 3}, 
	{4, 2, 6, 8, 5, 3, 7, 9, 1}, 
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

func TestSolveSudoku(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/solve-sudoku", solveSudoku)

	sendRequest := func(method string, path string, body interface{}) *httptest.ResponseRecorder {
		reqBody, _ := json.Marshal(body)
		req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		return w
	}

	// Test case: Valid Sudoku input
	t.Run("Valid Sudoku", func(t *testing.T) {
		w := sendRequest("POST", "/solve-sudoku", sudokuDto{Sudoku: validSudokuMock})

		assert.Equal(t, http.StatusOK, w.Code)

		var response sudokuDto
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, solvedSudokuMock, response.Sudoku)
	})

	// Test case: Invalid Sudoku input (bad JSON)
	t.Run("Invalid JSON", func(t *testing.T) {
		w := sendRequest("POST", "/solve-sudoku", `invalid json`)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Request is not valid")
	})

	// Test case: Invalid Sudoku grid
	t.Run("Invalid Sudoku", func(t *testing.T) {
		invalidSudoku := [][]int{
			{5, 5, 5, 5, 5, 5, 5, 5, 5},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
		}
		w := sendRequest("POST", "/solve-sudoku", sudokuDto{Sudoku: invalidSudoku})

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Sudoku is not valid")
	})
}