package cli

func (s *SimulatorCli) initGrid() {
	s.grid.
		SetRows(0).
		SetColumns(30, 0, 50).
		SetBorders(true).
		AddItem(s.menu, 0, 0, 1, 1, 0, 0, false).
		// AddItem(s.userGuidePage, 0, 1, 1, 1, 0, 0, false).
		AddItem(s.setProgramPage, 0, 1, 1, 1, 0, 0, false).
		AddItem(s.console, 0, 2, 1, 1, 0, 0, false)
}
