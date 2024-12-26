package cli

import "github.com/rivo/tview"

func (s *SimulatorCli) initMenu() {
	s.menu.
		SetDirection(tview.FlexRow).
		AddItem(s.button1, 1, 1, true).
		AddItem(nil, 1, 1, false).
		AddItem(s.button2, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(s.button3, 1, 1, false)
}
