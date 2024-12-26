package cli

func (s *SimulatorCli) initGrid() {
	s.grid.
		SetRows(0).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(s.menu, 0, 0, 1, 1, 0, 0, false).
		AddItem(s.userGuide, 0, 1, 1, 1, 0, 0, false)
}
