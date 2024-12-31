package cli

import "github.com/rivo/tview"

const UserGuide = `This is a simulator for the IAS computer.
Instructions:
1. Click the init button and initialize your program.
   You can use the load button at the bottom of the page to import the yaml files in the example folder.
   After completing the configuration, don’t forget to click the apply button to apply the settings to the simulator.
   If you can't see the button, the window may be too small. Please adjust it.
2. Click the run button to switch to the run page.
   Here you can see the data of each register
   You can choose to run directly or perform single-step debugging
`

func (s *SimulatorCli) initUserGuidePage() {
	s.userGuidePage = tview.NewTextView().SetTextAlign(tview.AlignLeft)
	s.userGuidePage.SetText(UserGuide)
}

func (s *SimulatorCli) navigateToUserGuidePage() {
	s.ClearMainGrid()
	s.mainGrid.AddItem(s.userGuidePage, 0, 1, 1, 1, 0, 0, false)
}
