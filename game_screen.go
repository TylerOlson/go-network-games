package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type GameScreen struct {
	currentGame *TicTacToeGame
	infoText    string
	*ScreenManager
	*TerminalScreen
}

// Create our very useful GameScreen
func NewGameScreen(screenManager *ScreenManager, ts *TerminalScreen) *GameScreen {
	return &GameScreen{
		infoText:       "Waiting for X to place",
		ScreenManager:  screenManager,
		TerminalScreen: ts,
	}
}

func (gs *GameScreen) ScreenStarted() {
	gs.infoText = "Waiting for X to place"
	gs.currentGame = NewTicTacToeGame(gs)
}

// Our very useful DrawContent func
func (gs *GameScreen) DrawContent() {
	if gs.currentGame.winner != ' ' {
		gs.infoText = fmt.Sprintf("Winner %c. New game? (y/n)", gs.currentGame.winner)
	}
	if gs.currentGame.winner == 'n' {
		gs.infoText = "Tie. New game? (y/n)"
	}
	drawText(gs.s, (gs.width/2)-(len(gs.infoText)/2), 9, (gs.width/2)+(len(gs.infoText)/2)+1, 9, gs.textStyle, gs.infoText) // Draw info text

	gs.currentGame.PrintBoard()
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
		case 110: //n
			gs.SetCurrentScreen(0) // +1 because main menu is 0
		case 121: //y
			if gs.currentGame.winner != ' ' {
				gs.ScreenStarted()
			}
		}

		if moveNum != 0 {
			gs.currentGame.DoMove(moveNum)
		}

	}

}
