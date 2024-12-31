package main

import (
	"github.com/caoqiming/ias-simulator/cli"
	"github.com/caoqiming/ias-simulator/simulator"
)

func main() {
	simulator.Init()
	cli.Init()
	cli.SimulatorCliSingleton.Run()
}
