package simulator

type IasSimulator struct {
}

// TakeNextInstruction take the next instruction into IR
func (*IasSimulator) TakeNextInstruction() {
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

func Init() {
	initMemory()
	initRegister()
	initFlag()
}
