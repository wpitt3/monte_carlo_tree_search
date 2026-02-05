package mcts

import "fmt"

type TicTacToe struct {
	board [3][3]int
}

type TicTacToeAction struct {
	x int
	y int
}

func (board TicTacToe) Copy() State[Action] {
	var newBoard TicTacToe
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard.board[i][j] = board.board[i][j]
		}
	}
	return newBoard
}

func (board TicTacToe) ValidActions() []Action {
	validMoves := make([]Action, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.board[i][j] == 0 {
				validMoves = append(validMoves, TicTacToeAction{i, j})
			}
		}
	}
	return validMoves
}

func (board TicTacToe) IsEndState() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (board TicTacToe) Winner() int {
	for i := 0; i < 3; i++ {
		var sum = 0
		for j := 0; j < 3; j++ {
			sum += board.board[i][j]
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	for i := 0; i < 3; i++ {
		var sum = 0
		for j := 0; j < 3; j++ {
			sum += board.board[j][i]
		}
		if abs(sum) == 3 {
			return sum / 3
		}
	}
	var sum = board.board[0][0] + board.board[1][1] + board.board[2][2]
	if abs(sum) == 3 {
		return sum / 3
	}
	sum = board.board[2][0] + board.board[1][1] + board.board[0][2]
	if abs(sum) == 3 {
		return sum / 3
	}
	return 0
}

func (board TicTacToe) PerformMove(action Action) State[Action] {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.board[j][i] = -1 * board.board[j][i]
		}
	}

	ticTacToeAction := action.(TicTacToeAction)
	board.board[ticTacToeAction.x][ticTacToeAction.y] = -1
	return board
}

func (board TicTacToe) PrintState() {
	valueToPrint := map[int]string{0: " ", 1: "#", -1: "O"}
	var toPrint = ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			toPrint += valueToPrint[board.board[j][i]]
		}
		toPrint += "\n"
	}
	fmt.Println(toPrint)
}

func (board TicTacToe) GetState() [][]int {
	slice := make([][]int, len(board.board))
	for i := range board.board {
		slice[i] = board.board[i][:]
	}
	return slice
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
