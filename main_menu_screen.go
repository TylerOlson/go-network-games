package main

import (
	"github.com/gdamore/tcell/v2"
)

// Define our MainMenuScreen, it embeds our TerminalScreen with a few more important variables
type MainMenuScreen struct {
	title            string
	currentSelection int
	options          []string
	highlight_style  tcell.Style
	*ScreenManager
	*TerminalScreen
}

// Create a new MainMenuScreen and point to it with list of options
func NewMainMenuScreen(ts *TerminalScreen, sm *ScreenManager, options ...string) *MainMenuScreen {
	return &MainMenuScreen{
		title:            "Main Menu",
		options:          options,
		currentSelection: 0,
		highlight_style:  tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite),
		TerminalScreen:   ts,
		ScreenManager:    sm,
	}
}

func (m *MainMenuScreen) ScreenStarted() {}

// Our own DrawContent func, prints our options defined before and our title
func (m *MainMenuScreen) DrawContent() {
	drawText(m.s, (m.width/2)-(len(m.title)/2), 2, (m.width/2)+(len(m.title)/2)+1, 2, m.textStyle, m.title)

	for i, option := range m.options {
		style := m.textStyle
		if m.currentSelection == i {
			style = m.highlight_style
		}
		drawText(m.s, 0, 6+i, len(option), 6+i, style, option)
	}

}

// Our OnKeyEvent, this is where we write our menu logic
func (m *MainMenuScreen) OnKeyEvent(key tcell.Key, ch rune) {
	if (key == 256 && ch == 119) || (key == 257 && ch == 0) { //w/up
		m.currentSelection--
	} else if (key == 256 && ch == 115) || (key == 258 && ch == 0) { //s/down
		m.currentSelection++
	} else if key == 13 && ch == 13 { //enter
		if m.currentSelection == len(m.options)-1 {
			Quit(m.s)
		}
		m.SetCurrentScreen(m.currentSelection + 1) // +1 because main menu is 0
	}

	if m.currentSelection >= len(m.options) {
		m.currentSelection = 0
	} else if m.currentSelection < 0 {
		m.currentSelection = len(m.options) - 1
	}
}
