package cli

import "github.com/rivo/tview"

type SimulatorCli struct {
	app       *tview.Application
	grid      *tview.Grid
	userGuide *tview.TextView
	menu      *tview.Flex
	button1   *tview.Button // 加载程序
	button2   *tview.Button // 开始运行
	button3   *tview.Button // 保存
}

var SimulatorCliSingleton *SimulatorCli

func init() {
	SimulatorCliSingleton = &SimulatorCli{
		app:       tview.NewApplication(),
		grid:      tview.NewGrid(),
		userGuide: tview.NewTextView().SetTextAlign(tview.AlignLeft),
		menu:      tview.NewFlex(),
	}
	SimulatorCliSingleton.initUserGuide()
	SimulatorCliSingleton.initGrid()
	SimulatorCliSingleton.initButton()
	SimulatorCliSingleton.initMenu()
}

func (s *SimulatorCli) Run() {
	if err := s.app.SetRoot(s.grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
