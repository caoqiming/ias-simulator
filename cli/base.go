package cli

import (
	"github.com/caoqiming/ias-simulator/simulator"
	"github.com/rivo/tview"
)

type SimulatorCli struct {
	program        ProgramConfig
	app            *tview.Application
	grid           *tview.Grid
	userGuidePage  *tview.TextView
	setProgramPage *tview.Form
	menu           *tview.Flex
	button1        *tview.Button // 加载程序
	button2        *tview.Button // 开始运行
	button3        *tview.Button // 保存

	console *tview.TextView // for debug
}

var SimulatorCliSingleton *SimulatorCli

func init() {
	SimulatorCliSingleton = &SimulatorCli{
		app:            tview.NewApplication(),
		grid:           tview.NewGrid(),
		setProgramPage: tview.NewForm(),
		menu:           tview.NewFlex(),
		console:        tview.NewTextView(),
	}
	SimulatorCliSingleton.program.MaxSteps = simulator.DefaultMaxSteps
	SimulatorCliSingleton.program.HaltAt = simulator.DefaultHaultAt

	SimulatorCliSingleton.initUserGuidePage()
	SimulatorCliSingleton.initSetProgramPage()
	SimulatorCliSingleton.initGrid()
	SimulatorCliSingleton.initButton()
	SimulatorCliSingleton.initMenu()

}

func (s *SimulatorCli) Run() {
	if err := s.app.SetRoot(s.grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
