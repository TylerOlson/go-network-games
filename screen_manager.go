package main

type ScreenManager struct {
	currentScreen TerminalScreenInterface
	screens       []TerminalScreenInterface
}

func NewScreenManager(screens ...TerminalScreenInterface) *ScreenManager {
	return &ScreenManager{}
}

func (sm *ScreenManager) SetCurrentScreen(newScreen int) {
	sm.currentScreen = sm.screens[newScreen]
}

func (sm *ScreenManager) AddScreen(newScreen TerminalScreenInterface) {
	sm.screens = append(sm.screens, newScreen)
}
