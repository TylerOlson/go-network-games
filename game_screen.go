package main

import (
	"github.com/gdamore/tcell/v2"
)

type GameScreen struct {
	currentGame *TicTacToeGame
	isSolo      bool
	infoText    string
	*TerminalScreen
}

// Create our very useful GameScreen
func NewGameScreen(isSolo bool, ts *TerminalScreen) *GameScreen {
	return &GameScreen{
		isSolo:         isSolo,
		infoText:       "Waiting for X to place",
		TerminalScreen: ts,
	}
}

func (gs *GameScreen) ScreenStarted() {
	gs.currentGame = NewTicTacToeGame()

	if gs.isSolo {
		gs.currentGame.player1 = NewTicTacToePlayer(true)
		gs.currentGame.player2 = NewTicTacToePlayer(false)
	}
}

// Our very useful DrawContent func
func (gs *GameScreen) DrawContent() {
	drawText(gs.s, (gs.width/2)-(len(gs.infoText)/2), 9, (gs.width/2)+(len(gs.infoText)/2)+1, 9, gs.textStyle, gs.infoText) // draw board value

	gs.currentGame.PrintBoard(gs)
}

func (gs *GameScreen) OnKeyEvent(key tcell.Key, ch rune) {
	if key == 256 {
		moveNum := 0
		switch ch {
		case 49: // 1
			moveNum = 1
		case 50: // 2
			moveNum = 2
		case 51: // 3
			moveNum = 3
		case 52: // 4
			moveNum = 4
		case 53: // 5
			moveNum = 5
		case 54: //6
			moveNum = 6
		case 55: // 7
			moveNum = 7
		case 56: // 8
			moveNum = 8
		case 57: // 9
			moveNum = 9
		}

		gs.currentGame.DoMove(moveNum, gs)
	}

}
