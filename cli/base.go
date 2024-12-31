package cli

import (
	"github.com/caoqiming/ias-simulator/simulator"
	"github.com/rivo/tview"
)

type SimulatorCli struct {
	program        ProgramConfig
	app            *tview.Application
	mainGrid       *tview.Grid
	userGuidePage  *tview.TextView
	setProgramPage *tview.Form
	runProgramPage *tview.Grid
	machineStatus  *tview.TextView

	menu                   *tview.Flex
	buttonToUserGuidePage  *tview.Button
	buttonToSetProgramPage *tview.Button
	buttonToRunProgramPage *tview.Button
	buttonRun              *tview.Button
	buttonRunSingleStep    *tview.Button

	console *tview.TextView // for debug
}

var SimulatorCliSingleton *SimulatorCli

func Init() {
	SimulatorCliSingleton = &SimulatorCli{
		app:            tview.NewApplication(),
		mainGrid:       tview.NewGrid(),
		setProgramPage: tview.NewForm(),
		runProgramPage: tview.NewGrid(),
		machineStatus:  tview.NewTextView(),
		menu:           tview.NewFlex(),
		console:        tview.NewTextView(),
	}
	SimulatorCliSingleton.program.MaxSteps = simulator.DefaultMaxSteps
	SimulatorCliSingleton.program.HaltAt = simulator.DefaultHaultAt

	SimulatorCliSingleton.initButton()
	SimulatorCliSingleton.initMenu()
	SimulatorCliSingleton.initUserGuidePage()
	SimulatorCliSingleton.initSetProgramPage()
	SimulatorCliSingleton.initRunProgramPage()
	SimulatorCliSingleton.initGrid()

}

func (s *SimulatorCli) Run() {
	if err := s.app.SetRoot(s.mainGrid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
