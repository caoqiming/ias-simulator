package cli

import "github.com/rivo/tview"

// 菜单栏的按钮
func (s *SimulatorCli) initButton() {
	s.buttonToUserGuidePage = tview.NewButton("help").SetSelectedFunc(s.navigateToUserGuidePage)

	s.buttonToSetProgramPage = tview.NewButton("init").SetSelectedFunc(s.navigateToSetProgramPage)

	s.buttonToRunProgramPage = tview.NewButton("run").SetSelectedFunc(s.navigateToRunProgramPage)

	s.buttonRun = tview.NewButton("run").SetSelectedFunc(s.runProgram)

	s.buttonRunSingleStep = tview.NewButton("single step").SetSelectedFunc(s.runProgramSingaleStep)
}
