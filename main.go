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

	w, h := s.Size()

	// Create a TerminalScreen for all of our other screens to look at
	ts := NewTerminalScreen(s, w, h, textStyle)          // typeof *TerminalScreen
	mms := NewMainMenuScreen(ts, "Option 1", "Option 2") // typeof *MainMenuScreen
	gs := NewGameScreen(ts)                              // typeof *GameScreen

	/*
		MainMenuScreen {
			title            string
			currentSelection int
			options          []string
			*TerminalScreen {
				s             tcell.Screen
				width, height int
				textStyle     tcell.Style
			}
		}

	*/

	// We create an empty currentScreen variable with type TerminalScreenInterface
	var currentScreen TerminalScreenInterface
	/*
		TerminalScreenInterface {
			DrawContent()
			OnKeyEvent(key tcell.Key, ch rune)
			UpdateSize(width, height int)
		}
	*/

	// Since type TerminalScreenInterface requires 3 functions, we can set currentScreen to mms,
	// as mms fulfills the requirement of those 3 functions. It doesn't matter where they are defined
	// as they are declared by the interface
	currentScreen = mms
	// so it's worth noting that here, currentScreen IS NOT of type MainMenuScreen. It is still of type
	// TerminalScreenInterface. Therefore it can ONLY do currentScreen.DrawContent()... and the other two funcs

	// the main confusing part about the interfaces is that they are implemented implicitly
	// so nothing really explicitly says that mms can be a TerminalScreenInterface, you just gotta know

	//_ = mms.title ✅
	//_ = currentScreen.title ❌

	/*
		type MainMenuScreen struct {
		title            string
		currentSelection int
		options          []string
		*TerminalScreen
	}*/

	//Quit function
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		s.Clear()

		// Draw our currentScreen
		currentScreen.DrawContent()

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
			currentScreen.OnKeyEvent(key, ch)

			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
			if ch == 119 { //w
				currentScreen = mms
			} else if ch == 115 { //s
				currentScreen = gs
			}
		}
	}

}
