package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// Define our MainMenuScreen, it embeds our TerminalScreen with a few more important variables
type MainMenuScreen struct {
	title            string
	currentSelection int
	options          []string
	*TerminalScreen
}

// Create a new MainMenuScreen and point to it with list of options
func NewMainMenuScreen(ts *TerminalScreen, options ...string) *MainMenuScreen {
	return &MainMenuScreen{
		options:          options,
		currentSelection: 0,
		TerminalScreen:   ts,
	}
}

// Our own DrawContent func, prints our options defined before and our title
func (m *MainMenuScreen) DrawContent() {
	drawText(m.s, (m.width/2)-(len(m.title)/2), 2, (m.width/2)+(len(m.title)/2)+1, 2, m.textStyle, m.title)

	for i, option := range m.options {
		drawText(m.s, 0, 6+i, len(option), 6+i, m.textStyle, option)
	}

	option1 := fmt.Sprintf("MainMenu %d %d", m.width, m.height)
	drawText(m.s, 0, 5, len(option1), 5, m.textStyle, option1)

}

// Our OnKeyEvent, this is where we write our menu logic
func (m *MainMenuScreen) OnKeyEvent(key tcell.Key, ch rune) {
	if ch == 119 { //w
		m.currentSelection--
	} else if ch == 115 { //s
		m.currentSelection++
	}

	if m.currentSelection > len(m.options) {
		m.currentSelection = 0
	} else if m.currentSelection < 0 {
		m.currentSelection = len(m.options) - 1
	}
}
