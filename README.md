# Sudoku Solver Service
This is a simple Sudoku solver web application built with the Gin framework in Go. The application provides a RESTful API to solve Sudoku puzzles. You can submit a Sudoku puzzle via a POST request, and the application will return the solved puzzle if it is solvable.

## Run the Application
Run the application with the `go run main.go` command.

## Run Tests
### Run Tests in `main` package
Run the tests in the `main` package with the `go test` command.

### Run Tests in `sudokusolver` package
Run the tests in the `sudokusolver` package with the `go test ./sudokusolver` command.

## API Usage
`POST /solve-sudoku`  
This endpoint accepts a JSON payload containing a 9x9 Sudoku table and returns the solved puzzle if it is solvable.

### Request
- **URL:** `/solve-sudoku`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body:** A JSON object with a single key `sudoku` that holds a 9x9 2D array of integers
- **Example Request Body:**  
```json
{
  "sudoku": [
    [5, 3, 0, 0, 7, 0, 0, 0, 0],
    [6, 0, 0, 1, 9, 5, 0, 0, 0],
    [0, 9, 8, 0, 0, 0, 0, 6, 0],
    [8, 0, 0, 0, 6, 0, 0, 0, 3],
    [4, 0, 0, 8, 0, 3, 0, 0, 1],
    [7, 0, 0, 0, 2, 0, 0, 0, 6],
    [0, 6, 0, 0, 0, 0, 2, 8, 0],
    [0, 0, 0, 4, 1, 9, 0, 0, 5],
    [0, 0, 0, 0, 8, 0, 0, 7, 9]
  ]
}
```

### Response
- **Success (200 OK):** The solved Sudoku puzzle is return in the same JSON format as in the request body
- **Failure (400 Bad Request):** Returns an error message if the input is invalid or the puzzle cannot be solved
- **Example Success Response:**  
```json
{
  "sudoku": [
    [5, 3, 4, 6, 7, 8, 9, 1, 2],
    [6, 7, 2, 1, 9, 5, 3, 4, 8],
    [1, 9, 8, 3, 4, 2, 5, 6, 7],
    [8, 5, 9, 7, 6, 1, 4, 2, 3],
    [4, 2, 6, 8, 5, 3, 7, 9, 1],
    [7, 1, 3, 9, 2, 4, 8, 5, 6],
    [9, 6, 1, 5, 3, 7, 2, 8, 4],
    [2, 8, 7, 4, 1, 9, 6, 3, 5],
    [3, 4, 5, 2, 8, 6, 1, 7, 9]
  ]
}
```

- **Example Failure Responses**:
1. **Invalid Sudoku:**  
```json
{
  "message": "Sudoku is not valid"
}
```

2. **The application could not solve the Sudoku:**  
```json
{
  "message": "Sudoku cannot be solved"
}
```

## Technologies Used
- [Go](https://go.dev/) v1.22.5
- [Gin Web Framework](https://gin-gonic.com/) v1.10.0
- [Testify](https://github.com/stretchr/testify) v1.9.0