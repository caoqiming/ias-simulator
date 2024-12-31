package cli

func (s *SimulatorCli) initGrid() {
	s.mainGrid.
		SetRows(0).
		SetColumns(10, 0, 40).
		SetBorders(true).
		AddItem(s.menu, 0, 0, 1, 1, 0, 0, false).
		AddItem(s.userGuidePage, 0, 1, 1, 1, 0, 0, false).
		AddItem(s.console, 0, 2, 1, 1, 0, 0, false)
}

// 清除当前组件，不清除会导无法正常focus。因为grid add item 时并不会清理同一个位置上的旧item，而是堆叠上去
func (s *SimulatorCli) ClearMainGrid() {
	s.mainGrid.RemoveItem(s.userGuidePage).
		RemoveItem(s.setProgramPage).
		RemoveItem(s.runProgramPage)
}
