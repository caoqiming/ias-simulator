package cli

import "github.com/rivo/tview"

func (s *SimulatorCli) initMenu() {
	s.menu.
		SetDirection(tview.FlexRow).
		AddItem(s.buttonToSetProgramPage, 1, 1, true).
		AddItem(nil, 1, 1, false).
		AddItem(s.buttonToRunProgramPage, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(s.buttonToUserGuidePage, 1, 1, false)

}
