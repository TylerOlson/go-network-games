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

	// Set text stylesÍ
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	w, h := s.Size()

	screenManager := NewScreenManager()

	// Create a TerminalScreen for all of our other screens to look at
	ts := NewTerminalScreen(s, w, h, defStyle)                        // typeof *TerminalScreen
	mms := NewMainMenuScreen(ts, screenManager, "Start Game", "Exit") // typeof *MainMenuScreen
	screenManager.AddScreen(mms)

	gs := NewGameScreen(ts) // typeof *GameScreen
	screenManager.AddScreen(gs)

	screenManager.SetCurrentScreen(0)

	text := ""
	//Quit function

	for {
		// Update screen
		s.Clear()

		// Draw our currentScreen
		screenManager.currentScreen.DrawContent()

		drawText(s, 0, 0, len(text), 0, defStyle, text)

		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			w, h = s.Size()
			ts.UpdateSize(w, h)
			s.Sync()
		case *tcell.EventKey:
			key, ch := ev.Key(), ev.Rune()
			screenManager.currentScreen.OnKeyEvent(key, ch)
			text = fmt.Sprintf("key %d ch %d", key, ch)

			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				Quit(s)
			}
		}
	}

}

func Quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}
