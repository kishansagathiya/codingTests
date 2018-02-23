package main

/*
use `go run tictactoe.go` to run it
*/

import (
	"fmt"
)

func main() {
	fmt.Print("This is a code for Tic Tac Toe\n")

	var board [3][3]int
	end := false
	var x int
	var y int
	left := 9

	for !end && left > 0 {
		turn := 1
		for turn < 3 {
			fmt.Printf("Player %d: Make your move(x coordinate, y coordinatem and your move 1 if 0, 2 if X)\n", turn)
			_, _ = fmt.Scanf("%d", &x)
			_, _ = fmt.Scanf("%d", &y)

			if board[x][y] == 0 {
				board[x][y] = turn
				left--
				turn++
				printBoard(board)
				if checkIfWon(board) {
					fmt.Printf("Player %d has won\n", turn)
					end = true
					break
				}
			} else {
				fmt.Print("invalid Move")
			}
		}
	}
}

func checkIfWon(board [3][3]int) bool {

	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][1] != 0 && board[i][2] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
		if board[0][i] != 0 && board[1][i] != 0 && board[2][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}

	if board[0][0] != 0 && board[1][1] != 0 && board[2][2] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}

	if board[0][2] != 0 && board[1][1] != 0 && board[2][0] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}

	return false
}

func printBoard(board [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Printf("\n")
	}
}
