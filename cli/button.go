package cli

import "github.com/rivo/tview"

func (s *SimulatorCli) initButton() {
	s.button1 = tview.NewButton("init")
	s.button2 = tview.NewButton("run")
	s.button3 = tview.NewButton("save")
}
