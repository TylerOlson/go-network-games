package main

import (
	"github.com/gdamore/tcell/v2"
)

// TerminalScreenInterface declares what functions are required by every screen we create
type TerminalScreenInterface interface {
	ScreenStarted()
	DrawContent()
	OnKeyEvent(key tcell.Key, ch rune)
	UpdateSize(width, height int)
}

// TerminalScreen struct, every screen needs these variables
type TerminalScreen struct {
	s             tcell.Screen
	width, height int
	textStyle     tcell.Style //actually doesnt need this but we move
}

// Create new TerminalScreen and point to it
func NewTerminalScreen(s tcell.Screen, width, height int, textStyle tcell.Style) *TerminalScreen {
	return &TerminalScreen{s, width, height, textStyle}
}

// Update size variables when resized, we can make it here because every screen needs it and doesn't change it
func (ts *TerminalScreen) UpdateSize(width, height int) {
	ts.width = width
	ts.height = height
}
