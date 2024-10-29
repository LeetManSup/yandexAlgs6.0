package main

import (
	"fmt"
	"slices"
)

func readBoard() [][]rune {
	var n byte
	fmt.Scan(&n)
	result := make([][]rune, n)

	var tmp string
	for i := byte(0); i < n; i++ {
		fmt.Scan(&tmp)
		result[i] = []rune(tmp)
	}
	return result
}

func processBoard(board [][]rune) [][]rune {
	isBoardEmpty := func(board [][]rune) bool {
		hasSharp := false

		for _, row := range board {
			if slices.Contains(row, '#') {
				hasSharp = true
				break
			}
		}
		return !hasSharp
	}

	strip := func(board [][]rune) [][]rune {
		isEdge := true
		for isEdge {
			if !slices.Contains(board[0], '#') && len(board) != 1 {
				board = board[1:]
			} else {
				isEdge = false
			}
		}

		isEdge = true
		for isEdge {
			if !slices.Contains(board[len(board)-1], '#') {
				board = board[:len(board)-1]
			} else {
				isEdge = false
			}
		}
		return board
	}

	transpose := func(board [][]rune) [][]rune {
		resultBoard := make([][]rune, len(board[0]))
		for _, row := range board {
			for j, cell := range row {
				resultBoard[j] = append(resultBoard[j], cell)
			}
		}
		return resultBoard
	}

	compress := func(board [][]rune) [][]rune {
		for i := len(board) - 2; i >= 0; i-- {
			if slices.Equal(board[i], board[i+1]) {
				board = append(board[:i], board[i+1:]...)
			}
		}
		return board
	}

	if isBoardEmpty(board) {
		return [][]rune{{46}}
	}

	board = strip(board)
	board = compress(board)
	board = transpose(board)
	board = strip(board)
	board = compress(board)
	board = transpose(board)

	return board
}

func detectLetter(board [][]rune) string {
	equals := func(board1, board2 [][]rune) bool {
		if len(board1) != len(board2) {
			return false
		}
		for i := range board1 {
			if len(board1[i]) != len(board2[i]) {
				return false
			}
			for j := range board1[i] {
				if board1[i][j] != board2[i][j] {
					return false
				}
			}
		}
		return true
	}

	if equals(board, [][]rune{{35}}) {
		return "I"
	} else if equals(board, [][]rune{{35, 35, 35}, {35, 46, 35}, {35, 35, 35}}) {
		return "O"
	} else if equals(board, [][]rune{{35, 35}, {35, 46}, {35, 35}}) {
		return "C"
	} else if equals(board, [][]rune{{35, 46}, {35, 35}}) {
		return "L"
	} else if equals(board, [][]rune{{35, 46, 35}, {35, 35, 35}, {35, 46, 35}}) {
		return "H"
	} else if equals(board, [][]rune{{35, 35, 35}, {35, 46, 35}, {35, 35, 35}, {35, 46, 46}}) {
		return "P"
	} else {
		return "X"
	}
}

func main() {
	board := readBoard()
	board = processBoard(board)
	letter := detectLetter(board)
	fmt.Println(letter)
}
