package main

type TicTacToeGame struct {
	board [3][3]rune
}

func NewTicTacToeGame() *TicTacToeGame {
	game := &TicTacToeGame{}
	game.ClearBoard()
	return game
}

func (game *TicTacToeGame) ClearBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.board[i][j] = '_'
		}
	}
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
