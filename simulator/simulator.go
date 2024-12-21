package simulator

import "fmt"

type InstructionAndAddr struct {
	OpCode byte
	Addr   int
}

// 将程序写入内存，用于初始化的环节。将指令按顺序从指定地址开始写入内存。
func SetInstructions(data []*InstructionAndAddr, startAddr int) {
	addr := startAddr
	for i := range data {
		if i%2 != 0 {
			continue
		}

		w := NewWord()
		// 写入当前指令到左边
		w.data[0] = data[i].OpCode
		w.data[1] = byte(data[i].Addr >> 4)
		w.data[2] = byte(data[i].Addr&0b00001111) << 4
		if i+1 < len(data) {
			// 写入下个指令到右边
			higherByte, lowerByte := ConvertIntToTwoByte(data[i+1].Addr)
			w.data[2] += data[i+1].OpCode & 0b11110000 >> 4
			w.data[3] = data[i+1].OpCode & 0b00001111 << 4
			w.data[3] += higherByte
			w.data[4] = lowerByte
		}
		DirectWrite(addr, w)
		addr++
	}
}

// TakeNextInstruction take the next instruction into IR
func TakeNextInstruction() {
	if !FlagIsNextInstructionInIBR {
		// take instruction from memory
		MAR.SetAddr(PC.GetAddr())
		memory.Read()
		// Left instruction required?
		leftCode, leftAddr := GetLeftInstructionFromMBR()
		rightCode, rightAddr := GetRightInstructionFromMBR()
		if FlagLeftInsturctionRequired {
			// Save right instruction to IBR
			IBR.Write(rightCode, rightAddr)
			FlagIsNextInstructionInIBR = true
			// Load left instruction to IR and MAR
			IR.Write(leftCode)
			MAR.SetAddr(leftAddr)
			return
		} else {
			// Left instruction not required
			// Load right insturction to IR and MAR
			IR.Write(rightCode)
			MAR.SetAddr(rightAddr)
			PC.Increase()
			FlagLeftInsturctionRequired = true
			return
		}
	} else {
		// Load IBR to IR and MAR
		code, addr := IBR.Read()
		IR.Write(code)
		MAR.SetAddr(addr)
		FlagIsNextInstructionInIBR = false
		PC.Increase()
		return
	}
}

// 执行一次
func ExecuteOneInstruction() error {
	instruction, err := instructionSet.GetInstruction(IR.Read())
	if err != nil {
		return err
	}

	instruction.Run()
	return nil
}

type SimulateOption func(*ExecuteParam)

func WithMaxSteps(maxSteps int) SimulateOption {
	return func(param *ExecuteParam) {
		param.maxSteps = maxSteps
	}
}

// 程序执行到该地址就结束（该地址的指令不会被执行）
func WithProgramExitAddr(programExitAddr int) SimulateOption {
	return func(param *ExecuteParam) {
		param.programExitAddr = programExitAddr
	}
}

type ExecuteParam struct {
	maxSteps        int
	programExitAddr int
}

// 连续执行，直到达到最大执行次数或遇到报错
func ExecuteWithMaxSteps(options ...SimulateOption) error {
	param := &ExecuteParam{
		maxSteps:        10000,
		programExitAddr: -1,
	}
	for _, option := range options {
		option(param)
	}

	for i := 0; i < param.maxSteps; i++ {
		TakeNextInstruction()
		if err := ExecuteOneInstruction(); err != nil {
			return fmt.Errorf("fail when execute step %d (start from 0), err: %v", i, err)
		}
		if PC.GetAddr() == param.programExitAddr {
			break
		}
	}

	return nil
}

func Init() {
	initMemory()
	initRegister()
	initFlag()
	InitInstructionSet()
}
