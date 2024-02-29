package main

import (
	"fmt"
	"math/rand"
)

const (
	Black = 0
	White = 1
)

type Position struct {
	x int
	y int
}

type Stone struct {
	color Color
	position Position
}

func main() {
	// Initialize the board
	board := make([]Stone, 9*9)
	for i := range board {
		board[i] = Stone{Black, Position{x: i%3 + 1, y: i/3 + 1}}
	}

	// Print the initial state of the board
	fmt.Println("Initial state of the board:")
	printBoard(board)

	// Play a random move
	move := rand.Intn(9*9)
	fmt.Printf("Playing random move %d...\n", move+1)
	playMove(board, move)

	// Print the final state of the board
	fmt.Println("Final state of the board:")
	printBoard(board)
}

func printBoard(board []Stone) {
	for i := range board {
		fmt.Printf("%d ", board[i].color)
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func playMove(board []Stone, move int) {
	// Check if the move is valid
	if !isValidMove(move) {
		fmt.Printf("Invalid move: %d\n", move+1)
		return
	}

	// Make the move
	board[move].color = White

	// Print the updated state of the board
	fmt.Println("Updated state of the board:")
	printBoard(board)
}

func isValidMove(move int) bool {
	return (move >= 0 && move < 9*9) && board[move].color == Black
}
