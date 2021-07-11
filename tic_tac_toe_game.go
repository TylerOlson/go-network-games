package main

import "fmt"

type TicTacToeGame struct {
	board          [3][3]rune
	rowX           [3]int
	columnX        [3]int
	diagX          [3]int
	diagOppX       [3]int
	rowO           [3]int
	columnO        [3]int
	diagO          [3]int
	diagOppO       [3]int
	moveCount      int
	winner         rune
	isCurrentMoveX bool
	gs             *GameScreen
}

func NewTicTacToeGame(gs *GameScreen) *TicTacToeGame {
	game := &TicTacToeGame{
		isCurrentMoveX: true,
		gs:             gs,
		winner:         ' ',
	}
	game.ClearBoard()
	return game
}

func (game *TicTacToeGame) ClearBoard() {
	num := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.board[i][j] = rune(num) + 48 // add to get ASCII value
			num++
		}
	}
}

func (game *TicTacToeGame) DoMove(moveNum int) {
	if game.winner != ' ' {
		return
	}

	num := 1
	moveRune := 'X'
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if num == moveNum {
				if !game.isCurrentMoveX {
					moveRune = 'O'
				}
				if game.board[i][j] == 'X' || game.board[i][j] == 'O' {
					game.gs.infoText = fmt.Sprintf("Cannot place %c at cell #%d, there is already an %c", moveRune, num, game.board[i][j])
					return
				}
				if game.isCurrentMoveX {
					game.board[i][j] = 'X'
					game.rowX[i]++
					game.columnX[j]++
					if i == j {
						game.diagX[i]++
					}
					if i+j+1 == 3 {
						game.diagOppX[i]++
					}

					moveRune = 'O'
				} else {
					game.board[i][j] = 'O'
					game.rowO[i]++
					game.columnO[j]++
					if i == j {
						game.diagO[i]++
					}
					if i+j+1 == 3 {
						game.diagOppO[i]++
					}

					moveRune = 'X'
				}
				game.moveCount++
				game.CheckWinner(i, j)
			}
			num++
		}
	}
	game.isCurrentMoveX = !game.isCurrentMoveX

	game.gs.infoText = fmt.Sprintf("Waiting for %c to place", moveRune)
}

func (game *TicTacToeGame) PrintBoard() {
	line := ""

	drawBox(game.gs.s, (game.gs.width/2)-11, 1, (game.gs.width/2)+11, 7, game.gs.textStyle, line) // Draw outline box

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			line += string(game.board[i][j]) // Add row to one line to print

			if j != 2 {
				line += "  |  " // Add seperator
			}

		}
		drawText(game.gs.s, (game.gs.width/2)-(len(line)/2), 2+i*2, (game.gs.width/2)+(len(line)/2)+1, 2+i*2, game.gs.textStyle, line) // Draw board value
		if i > 0 {
			line = "---------------------"
			drawText(game.gs.s, (game.gs.width/2)-(len(line)/2), 1+i*2, (game.gs.width/2)+(len(line)/2)+1, 1+i*2, game.gs.textStyle, line) // Draw divider
		}
		line = ""
	}

}

func (game *TicTacToeGame) CheckWinner(row, col int) {
	xCount, xOppCount, oCount, oOppCount := 0, 0, 0, 0
	for i := 0; i < 3; i++ { // Check sum of diag lists
		xCount += game.diagX[i]
		xOppCount += game.diagOppX[i]

		oCount += game.diagO[i]
		oOppCount += game.diagOppO[i]
	}

	if game.rowX[row] == 3 || game.columnX[col] == 3 || xCount == 3 || xOppCount == 3 {
		game.winner = 'X'
	}
	if game.rowO[row] == 3 || game.columnO[col] == 3 || oCount == 3 || oOppCount == 3 {
		game.winner = 'O'
	}

	if game.moveCount >= 9 {
		game.winner = 'n'
	}
}
