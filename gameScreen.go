package main

import "github.com/gdamore/tcell/v2"

type GameScreen struct {
	title         string
	s             tcell.Screen
	width, height int
	textStyle     tcell.Style
}

func newGameScreen(title string, s tcell.Screen, width, height int, textStyle tcell.Style) *GameScreen {
	gs := GameScreen{title, s, width, height, textStyle}
	return &gs
}

func (gs *GameScreen) DrawContent() {
	drawText(gs.s, (gs.width/2)-(len(gs.title)/2), 2, (gs.width/2)+(len(gs.title)/2)+1, 2, gs.textStyle, gs.title)
}

func (gs *GameScreen) UpdateSize(width, height int) {
	gs.width = width
	gs.height = height
}
