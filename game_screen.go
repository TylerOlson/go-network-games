package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type GameScreen struct {
	currentGame *TicTacToeGame
	*TerminalScreen
}

// Create our very useful GameScreen
func NewGameScreen(ts *TerminalScreen) *GameScreen {
	return &GameScreen{
		TerminalScreen: ts,
	}
}

func (gs *GameScreen) ScreenStarted() {
	gs.currentGame = NewTicTacToeGame()

}

// Our very useful DrawContent func
func (gs *GameScreen) DrawContent() {
	option1 := fmt.Sprintf("GameScreen %d %d", gs.width, gs.height)
	drawText(gs.s, 0, 5, len(option1), 5, gs.textStyle, option1)

	gs.currentGame.PrintBoard(gs)
}

func (gs *GameScreen) OnKeyEvent(key tcell.Key, ch rune) {}
