package main

import "fmt"

type TicTacToeGame struct {
	board          [3][3]rune
	player1        *TicTacToePlayer
	player2        *TicTacToePlayer
	isCurrentMoveX bool
}

func NewTicTacToeGame() *TicTacToeGame {
	game := &TicTacToeGame{}
	game.ClearBoard()
	game.isCurrentMoveX = true
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

func (game *TicTacToeGame) DoMove(moveNum int, gs *GameScreen) {
	// im lazy
	num := 1
	moveRune := 'X'
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if num == moveNum {
				if !game.isCurrentMoveX {
					moveRune = 'O'
				}
				if game.board[i][j] == 'X' || game.board[i][j] == 'O' {
					gs.infoText = fmt.Sprintf("Cannot place %c at cell #%d, there is already an %c", moveRune, num, game.board[i][j])
					return
				}
				if game.isCurrentMoveX {
					game.board[i][j] = 'X'
					moveRune = 'O'
				} else {
					game.board[i][j] = 'O'
					moveRune = 'X'
				}
			}
			num++
		}
	}
	game.isCurrentMoveX = !game.isCurrentMoveX

	gs.infoText = fmt.Sprintf("Waiting for %c to place", moveRune)
}

func (game *TicTacToeGame) PrintBoard(gs *GameScreen) {
	line := ""

	drawBox(gs.s, (gs.width/2)-11, 1, (gs.width/2)+11, 7, gs.textStyle, line)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			line += string(game.board[i][j])

			if j != 2 {
				line += "  |  "
			}
		}

		drawText(gs.s, (gs.width/2)-(len(line)/2), 2+i*2, (gs.width/2)+(len(line)/2)+1, 2+i*2, gs.textStyle, line) // draw board value
		if i > 0 {
			line = "---------------------"
			drawText(gs.s, (gs.width/2)-(len(line)/2), 1+i*2, (gs.width/2)+(len(line)/2)+1, 1+i*2, gs.textStyle, line) // draw divider
		}
		line = ""
	}

}

type TicTacToePlayer struct {
	isX bool
}

func NewTicTacToePlayer(isX bool) *TicTacToePlayer {
	return &TicTacToePlayer{isX: isX}
}
