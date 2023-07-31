package main

import "fmt"

func dfs(grid [][]int, row, col int) {
	rows := len(grid)
	cols := len(grid[0])

	// Check if the current cell is within the boundaries of the grid and it is unvisited (value is 1)
	if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == 1 {
		// Mark the current cell as visited by changing its value to 0
		grid[row][col] = 0

		// Recursively explore the adjacent cells in all four directions (up, down, left, right)
		dfs(grid, row+1, col) // Down
		dfs(grid, row-1, col) // Up
		dfs(grid, row, col+1) // Right
		dfs(grid, row, col-1) // Left
	}
}

func countRectangles(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	count := 0

	// Loop through each cell and count rectangles starting from unvisited 1's
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				count++
				dfs(matrix, i, j)
			}
		}
	}

	return count
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	rectanglesCount := countRectangles(matrix)
	fmt.Println("Number of rectangles formed by 1's with no adjacent rectangles:", rectanglesCount)
}
