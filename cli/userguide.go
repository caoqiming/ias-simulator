package cli

import "github.com/rivo/tview"

const UserGuide = `This is a simulator for the IAS computer.
Instructions:
1. Click the init button and initialize your program according to the prompts.
2. Click the run button to start running your program.
3. (optional) Click the save button to save your program.
`

func (s *SimulatorCli) initUserGuide() {
	s.userGuide = tview.NewTextView().SetTextAlign(tview.AlignLeft)
	s.userGuide.SetText(UserGuide)
}
