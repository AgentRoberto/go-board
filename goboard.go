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

func checkWin(board [][]string, player string) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}

	// Check diagonals
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}

	return false
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

		// Check for a win.
		if checkWin(board, currentPlayer) {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		// Switch to the other player.
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}

		moves++
	}

	if moves == 9 {
		fmt.Println("It's a draw! The game is over.")
	}
}
