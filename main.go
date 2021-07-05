package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set text styles
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	textStyle := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	w, h := s.Size()
	mainMenu := newGameScreen("Main Menu", s, w, h, textStyle)
	currentMenu := mainMenu

	text := ""

	//Quit function
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Clear()

		currentMenu.DrawContent()

		drawText(s, 1, 1, 20, 2, textStyle, text)

		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			w, h = s.Size()
			currentMenu.UpdateSize(w, h)
			s.Sync()
		case *tcell.EventKey:
			key, ch := ev.Key(), ev.Rune()
			text = fmt.Sprint("Key: ", key, " Rune: ", ch)
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Rune() == 119 { //w
				text = "W"
			} else if ev.Rune() == 115 { //s
				text = "S"
			}
		}
	}
}
