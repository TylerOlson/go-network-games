package main

import (
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

	s.SetContent(0, 0, 'H', nil, textStyle)
	s.SetContent(1, 0, 'i', nil, textStyle)
	s.SetContent(2, 0, '!', nil, textStyle)

	//Quit function
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}
