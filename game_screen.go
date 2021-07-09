package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type GameScreen struct {
	somethingUseful int
	*TerminalScreen
}

// Create our very useful GameScreen
func NewGameScreen(ts *TerminalScreen) *GameScreen {
	return &GameScreen{
		somethingUseful: 0,
		TerminalScreen:  ts,
	}
}

// Our very useful DrawContent func
func (gs *GameScreen) DrawContent() {
	option1 := fmt.Sprintf("GameScreen %d %d", gs.width, gs.height)
	drawText(gs.s, 0, 5, len(option1), 5, gs.textStyle, option1)
}

func (gs *GameScreen) OnKeyEvent(key tcell.Key, ch rune) {}
