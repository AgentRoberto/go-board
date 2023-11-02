package main

import (
	"fmt"
	"strings"
)

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, " "))
	}
}

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Print the initial board.
	printBoard(board)

	// Variable to keep track of the current player (X or O).
	currentPlayer := "X"

	// Number of moves played.
	moves := 0

	for moves < 9 {
		var row, col int

		// Prompt the user for input.
		fmt.Printf("Player %s, enter row (0, 1, 2) and column (0, 1, 2) separated by a space: ", currentPlayer)
		_, err := fmt.Scanf("%d %d", &row, &col)

		// Check for valid input.
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != "_" {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Update the board.
		board[row][col] = currentPlayer

		// Print the updated board.
		printBoard(board)

		// Check for a win condition (not implemented here).
		// You can add your win-checking logic here.

		// Switch to the other player.
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}

		moves++
	}

	fmt.Println("It's a draw! The game is over.")
}
