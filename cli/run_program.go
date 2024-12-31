package cli

import (
	"github.com/caoqiming/ias-simulator/simulator"
)

func (s *SimulatorCli) initRunProgramPage() {
	s.runProgramPage.Clear()
	// status,buttons
	s.runProgramPage.SetRows(5, 1, 1)
	// button,nil,button,empty
	s.runProgramPage.SetColumns(15, 1, 15, 0)
	s.runProgramPage.AddItem(s.machineStatus, 0, 0, 1, 4, 0, 0, false)
	s.runProgramPage.AddItem(s.buttonRun, 1, 0, 1, 1, 0, 0, true)
	s.runProgramPage.AddItem(s.buttonRunSingleStep, 1, 2, 1, 1, 0, 0, false)
	s.refreshMachineStatus()
}

func (s *SimulatorCli) navigateToRunProgramPage() {
	s.ClearMainGrid()
	s.mainGrid.AddItem(s.runProgramPage, 0, 1, 1, 1, 0, 0, false)
}

func (s *SimulatorCli) refreshMachineStatus() {
	s.machineStatus.SetText(simulator.SPrintStatus())
}

func (s *SimulatorCli) runProgram() {
	s.appendToConsole("start to run program")
	options := make([]simulator.SimulateOption, 0)
	if s.program.MaxSteps > 0 {
		options = append(options, simulator.WithMaxSteps(s.program.MaxSteps))
	}
	if s.program.HaltAt > 0 {
		options = append(options, simulator.WithHaultAt(s.program.HaltAt))
	}
	if err := simulator.Execute(options...); err != nil {
		s.appendToConsole(err.Error())
	}
	s.refreshMachineStatus()
}

func (s *SimulatorCli) runProgramSingaleStep() {
	s.appendToConsole("run single step")
	options := []simulator.SimulateOption{
		simulator.WithMaxSteps(1),
	}
	if s.program.HaltAt > 0 {
		options = append(options, simulator.WithHaultAt(s.program.HaltAt))
	}
	if err := simulator.Execute(options...); err != nil {
		s.appendToConsole(err.Error())
	}
	s.refreshMachineStatus()
}
