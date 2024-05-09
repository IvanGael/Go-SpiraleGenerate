package main

import (
	"fmt"
)

func generateSpiral(n int) {
	if n <= 0 {
		return
	}

	// Define directions: right, down, left, up
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Initialize variables
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	row, col := 0, 0
	dir := 0
	value := n * n

	for value > 0 {
		// Fill the current position
		matrix[row][col] = value
		value--

		// Calculate the next position
		nextRow := row + directions[dir][0]
		nextCol := col + directions[dir][1]

		// Check if the next position is valid and empty
		if nextRow >= 0 && nextRow < n && nextCol >= 0 && nextCol < n && matrix[nextRow][nextCol] == 0 {
			row, col = nextRow, nextCol
		} else {
			// Change direction
			dir = (dir + 1) % 4
			row += directions[dir][0]
			col += directions[dir][1]
		}
	}

	// Print the spiral matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	n := 8
	generateSpiral(n)
}
