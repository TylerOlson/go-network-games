package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "server" {
			println("Starting server")

			StartGameServer()
			return
		} else if args[0] == "client" {
			println("Starting client")

			StartClient("test")
			StartClient("d")

			return
		}
	}

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
	s.EnableMouse()
	s.Clear()

	w, h := s.Size()

	screenManager := NewScreenManager()

	// Create a TerminalScreen for all of our other screens to look at
	ts := NewTerminalScreen(s, w, h, defStyle)                             // typeof *TerminalScreen
	mms := NewMainMenuScreen(ts, screenManager, "Start Solo Game", "Exit") // typeof *MainMenuScreen
	gsSolo := NewGameScreen(screenManager, ts)                             // typeof *GameScreen
	//gsMulti := NewGameScreen(screenManager, ts)                                                      // typeof *GameScreen

	screenManager.AddScreen(mms)
	screenManager.AddScreen(gsSolo) // Solo tictactoe
	//screenManager.AddScreen(gsMulti) // Multiplayer

	screenManager.SetCurrentScreen(0)

	keyPressedText := "EventKey:"
	mouseText := "EventMouse:"

	for {
		// Update screen
		s.Clear()

		// Draw our currentScreen
		screenManager.currentScreen.DrawContent()

		drawText(s, 0, 0, len(keyPressedText), 0, defStyle, keyPressedText) // Draw key pressed
		drawText(s, 0, 1, len(mouseText), 1, defStyle, mouseText)           // Draw mouse info

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
			keyPressedText = fmt.Sprintf("EventKey: %d ch %d", key, ch)

			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				Quit(s)
			}
		case *tcell.EventMouse:
			mod := ev.Modifiers()
			buttons := ev.Buttons()
			// Only process button events, not wheel events
			buttons &= tcell.ButtonMask(0xff)

			x, y := ev.Position()

			screenManager.currentScreen.OnMouseEvent(mod, buttons, x, y)
			mouseText = fmt.Sprintf("EventMouse: %d Buttons: %d Position: %d,%d", mod, buttons, x, y)

		}
	}

}

//Quit function
func Quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}
