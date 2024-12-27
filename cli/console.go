package cli

import "fmt"

func (s *SimulatorCli) appendToConsole(text string) {
	s.console.SetText(fmt.Sprintf("%s\n%s", text, s.console.GetText(true)))
}
