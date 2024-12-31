package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/caoqiming/ias-simulator/simulator"
	"gopkg.in/yaml.v3"
)

const (
	SetProgramTips            = `You can type in a program or import a previously saved program.`
	ProgramAddrTips           = `Your program will be written to memory starting from this address.`
	ProgramCounterStartAtTips = `The program will start running from this address, that is, this is the initial value of PC.`
	MaxStepsTips              = `The program will stop after executing this number of instructions.`
	HaultAtTips               = `The program will stop when it reaches this address, and the instructions at this address will not be executed.`
	MemorySettingsTips        = `Specify the memory to be initialized here. 
	Because programs and data are stored in memory, you can actually write programs here, which is equivalent to writing programs on it. 
	You can specify multiple non-contiguous memory addresses to write here. Please follow the yaml format, please refer to README for details.`
)

// This page is used to write the program into memory,
// set the location to start running, and some settings during running.
func (s *SimulatorCli) initSetProgramPage() {
	s.setProgramPage.Clear(true)
	// program
	s.setProgramPage.AddTextView("", SetProgramTips, 0, 1, true, false)
	s.setProgramPage.AddTextArea("program", strings.Join(s.program.ProgramInHexFormat, "\n"), 0, 10, 0, func(text string) {
		programInHexFormat, err := splitAndValidateProgram(text)
		if err != nil {
			s.appendToConsole(err.Error())
			return
		}
		s.program.ProgramInHexFormat = programInHexFormat
		s.appendToConsole("program set successfully")
	})

	// where program will be written to
	s.setProgramPage.AddTextView("", ProgramAddrTips, 0, 1, true, false)
	s.setProgramPage.AddInputField("program addr", fmt.Sprintf("%d", s.program.ProgramAddr), 4, nil, func(text string) {
		if text == "" {
			return
		}
		result, err := strconv.Atoi(text)
		if err != nil {
			s.appendToConsole(fmt.Sprintf("program addr %s is not a valid integer", text))
		}
		s.program.ProgramAddr = result
	})

	// where program start
	s.setProgramPage.AddTextView("", ProgramCounterStartAtTips, 0, 1, true, false)
	s.setProgramPage.AddInputField("start at", fmt.Sprintf("%d", s.program.ProgramCounterStartAt), 4, nil, func(text string) {
		if text == "" {
			return
		}
		result, err := strconv.Atoi(text)
		if err != nil {
			s.appendToConsole(fmt.Sprintf("start at %s is not a valid integer", text))
		}
		s.program.ProgramCounterStartAt = result
	})

	// max steps
	s.setProgramPage.AddTextView("", MaxStepsTips, 0, 1, true, false)
	s.setProgramPage.AddInputField("max steps", fmt.Sprintf("%d", s.program.MaxSteps), 10, nil, func(text string) {
		if text == "" {
			return
		}
		result, err := strconv.Atoi(text)
		if err != nil {
			s.appendToConsole(fmt.Sprintf("max steps %s is not a valid integer", text))
		}
		s.program.MaxSteps = result
	})

	// hault at
	s.setProgramPage.AddTextView("", HaultAtTips, 0, 1, true, false)
	s.setProgramPage.AddInputField("hault at", fmt.Sprintf("%d", s.program.HaltAt), 4, nil, func(text string) {
		if text == "" {
			return
		}
		result, err := strconv.Atoi(text)
		if err != nil {
			s.appendToConsole(fmt.Sprintf("hault at %s is not a valid integer", text))
		}
		s.program.HaltAt = result
	})

	// memory init
	s.setProgramPage.AddTextView("", MemorySettingsTips, 0, 1, true, false)
	var memorySettingsStr string
	memorySettingsBytes, err := yaml.Marshal(s.program.MemorySettings)
	if err == nil {
		memorySettingsStr = string(memorySettingsBytes)
	} else {
		s.appendToConsole(fmt.Sprintf("fial to unmarshal memory settings from file, err: %s", err.Error()))
	}
	s.setProgramPage.AddTextArea("memory settings", memorySettingsStr, 0, 10, 0, func(text string) {
		if err := yaml.Unmarshal([]byte(text), &s.program.MemorySettings); err != nil {
			s.appendToConsole(fmt.Sprintf("fail to parse memory settings, error: %s", err.Error()))
		} else {
			s.appendToConsole("Memory data set successfully")
		}
	})

	// load/apply/save
	s.setProgramPage.AddInputField("load from/save to", s.program.ProgramPath, 0, nil, func(text string) {
		s.program.ProgramPath = text
	})
	s.setProgramPage.AddButton("load", s.loadPorgramFromPath)
	s.setProgramPage.AddButton("apply", s.applyProgramSettingsToSimulator)
	s.setProgramPage.AddButton("save", s.savePorgramToPath)
}

func (s *SimulatorCli) navigateToSetProgramPage() {
	s.ClearMainGrid()
	s.mainGrid.AddItem(s.setProgramPage, 0, 1, 1, 1, 0, 0, false)
}

func (s *SimulatorCli) loadPorgramFromPath() {
	s.appendToConsole(fmt.Sprintf("loadPorgramFromPath %s", s.program.ProgramPath))
	content, err := os.ReadFile(s.program.ProgramPath)
	if err != nil {
		s.appendToConsole(err.Error())
		return
	}

	if err = yaml.Unmarshal(content, &s.program); err != nil {
		s.appendToConsole(err.Error())
		return
	}

	// 刷新
	s.initSetProgramPage()
}

func (s *SimulatorCli) savePorgramToPath() {
	s.appendToConsole(fmt.Sprintf("savePorgramToPath %s", s.program.ProgramPath))
	content, err := yaml.Marshal(s.program)
	if err != nil {
		s.appendToConsole(err.Error())
		return
	}

	if err := os.WriteFile(s.program.ProgramPath, content, 0644); err != nil {
		s.appendToConsole(err.Error())
		return
	}
}

func (s *SimulatorCli) applyProgramSettingsToSimulator() {
	s.appendToConsole("apply program settings to simulator")
	// 清理当前的状态
	simulator.Init()
	// 写入程序
	for i, data := range s.program.ProgramInHexFormat {
		w, err := simulator.NewWordFromHexStr(data)
		if err != nil {
			s.appendToConsole(err.Error())
		}
		simulator.DirectWrite(s.program.ProgramAddr+i, w)
	}
	// 写入内存初始化
	for _, memorySetting := range s.program.MemorySettings {
		for i, data := range memorySetting.Content {
			w, err := simulator.NewWordFromHexStr(data)
			if err != nil {
				s.appendToConsole(err.Error())
			}
			simulator.DirectWrite(memorySetting.Addr+i, w)
		}
	}
	// set PC
	simulator.PC.SetAddr(s.program.ProgramCounterStartAt)
}
