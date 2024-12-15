package simulator

import "fmt"

const (
	// Data transfer
	// Transfer MQ to AC, MQ is multiply-quotient register
	OpcodeLoadMQ = 0b00001010
	// Transfer M(X) to MQ
	OpcodeLoadMToMQ = 0b00001001
	// Transfer AC to memory location x
	OpcodeStoreM = 0b00100001
	// Transfer M(X) to the AC, AC is accumulator register
	OpcodeLoadM = 0b00000001
	// Transfer -M(X) to AC
	OpcodeLoadNegativeM = 0b00000010
	// Transfer |M(X)| to AC
	OpcodeLoadAbsM = 0b00000011
	// Transfer -|M(X)| to AC
	OpcodeLoadNegativeAbsM = 0b00000100

	// Unconditional branch
	// Take next instructionfrom left half of M(X)
	OpcodeJumpMLeft = 0b00001101
	// Take next instructionfrom right half of M(X)
	OpcodeJumpMRight = 0b00001110

	// Conditional branch
	// If the number in the AC is nonnegative take the next instruction from left half of M(X)
	OpcodeConditionalJumpMLeft = 0b00001111
	// If the number in the AC is nonnegative take the next instruction from right half of M(X)
	OpcodeConditionalJumpMRight = 0b00010000

	// Arithmetic
	// Add M(X) to AC, put the result in AC
	OpcodeAddM = 0b00000101
	// Add |M(X)| to AC, put the result in AC
	OpcodeAddAbsM = 0b00000111
	// Substract M(X) from AC, put the resul in AC
	OpcodeSubstractM = 0b00000110
	// Substract |M(X)| from AC, put the resul in AC
	OpcodeSubstractAbsM = 0b00001000
	// Multiply M(X) by MQ, put most significant bits of result in AC, put least significant bits in MQ
	OpcodeMultiplyM = 0b00001011
	// Divide AC by M(X), put the quotient in MQ and the remainder in AC
	OpcodeDivideM = 0b00001100
	// Multiply AC by 2, that is, shift left one bit position
	OpcodeLSH = 0b00010100
	// Divide AC by 2, that is, shift right one position
	OpcodeRSH = 0b00010101

	// Addresss modify
	// Replace left address field at M(X)[8:19] by 12 rightmost bits of AC
	OpcodeStoreMLeft = 0b00010010
	// Replace right address field at M(X)[28:39] by 12 rightmost bits of AC
	OpcodeStoreMRight = 0b00010011
)

var instructionSet InstructionSet

type InstructionSet struct {
	instructions map[byte]InstructionInterface
}

func (is *InstructionSet) regist(code byte, instruction InstructionInterface) {
	is.instructions[code] = instruction
}
func (is *InstructionSet) GetInstruction(code byte) (InstructionInterface, error) {
	instruction, ok := is.instructions[code]
	if !ok {
		return nil, fmt.Errorf("instruction with op code %08b not found", code)
	}
	return instruction, nil
}

func InitInstructionSet() {
	instructionSet.regist(OpcodeLoadMQ, &InstructionLoadMQ{})
	instructionSet.regist(OpcodeLoadMToMQ, &InstructionLoadMToMQ{})
	instructionSet.regist(OpcodeStoreM, &InstructionStoreM{})
}

type InstructionInterface interface {
	Run()
}

// LOAD MQ
type InstructionLoadMQ struct{}

func (instruction *InstructionLoadMQ) Run() {
	AC.SetWord(MQ.GetWord())
}

// LOAD MQ,M(X)
type InstructionLoadMToMQ struct{}

func (instruction *InstructionLoadMToMQ) Run() {
	memory.Read()
	MQ.SetWord(MBR.GetWord())
}

// STOR M(X)

type InstructionStoreM struct{}

func (instruction *InstructionStoreM) Run() {
	MBR.SetWord(AC.GetWord())
	memory.Write()
}

// LOAD M(X)

type InstructionLoadM struct{}

func (instruction *InstructionLoadM) Run() {
	memory.Read()
	AC.SetWord(MBR.GetWord())
}

// LOAD -M(X)

type InstructionLoadNegativeM struct{}

func (instruction *InstructionLoadNegativeM) Run() {
	memory.Read()
	AC.SetWord(MBR.GetWord().Opposite())
}

// LOAD |M(X)|

type InstructionLoadAbsM struct{}

func (instruction *InstructionLoadAbsM) Run() {
	memory.Read()
	AC.SetWord(MBR.GetWord().Abs())
}

// LOAD -|M(X)|

type InstructionLoadNegativeAbsM struct{}

func (instruction *InstructionLoadNegativeAbsM) Run() {
	memory.Read()
	AC.SetWord(MBR.GetWord().Abs().Opposite())
}

// JUMP M(X,0:19)
type InstructionJumpMLeft struct{}

func (instruction *InstructionJumpMLeft) Run() {
	TakeInstructionFromLeftM()
}

// JUMP M(X,20:39)
type InstructionJumpMRight struct{}

func (instruction *InstructionJumpMRight) Run() {
	TakeInstructionFromRightM()
}

// JUMP +M(X,0:19)
type InstructionConditionalJumpMLeft struct{}

func (instruction *InstructionConditionalJumpMLeft) Run() {
	if !AC.IsNegative() {
		TakeInstructionFromLeftM()
	}
}

// JUMP +M(X,20:39)
type InstructionConditionalJumpMRight struct{}

func (instruction *InstructionConditionalJumpMRight) Run() {
	if !AC.IsNegative() {
		TakeInstructionFromRightM()
	}
}

// ADD M(X)
type InstructionAddM struct{}

func (instruction *InstructionAddM) Run() {
	memory.Read()
	r := AC.GetWord().Add(MBR.GetWord())
	AC.SetWord(r)
}

// ADD |M(X)|
type InstructionAddAbsM struct{}

func (instruction *InstructionAddAbsM) Run() {
	memory.Read()
	r := AC.GetWord().Add(MBR.GetWord().Abs())
	AC.SetWord(r)
}

// SUB M(X)
type InstructionSubM struct{}

func (instruction *InstructionSubM) Run() {
	memory.Read()
	r := AC.GetWord().Sub(MBR.GetWord())
	AC.SetWord(r)
}

// SUB |M(X)|
type InstructionSubAbsM struct{}

func (instruction *InstructionSubAbsM) Run() {
	memory.Read()
	r := AC.GetWord().Sub(MBR.GetWord().Abs())
	AC.SetWord(r)
}

// MUL M(x)
type InstructionMultiplyM struct{}

func (instruction *InstructionMultiplyM) Run() {
	// TODO
}

// helper function

func GetLeftInstructionFromMBR() (opcode byte, addr int) {
	// 0~7
	opcode = MBR.GetWord().data[0]
	// 8~19 12bit = 8bit+ 4bit
	addr = int(MBR.GetWord().data[1])<<4 + int(MBR.GetWord().data[2]&0b11110000)>>4
	return
}

func GetRightInstructionFromMBR() (opcode byte, addr int) {
	// 20~27
	opcode = MBR.GetWord().data[2]<<4 + MBR.GetWord().data[3]>>4
	// 28~39 12bit = 8bit+ 4bit
	addr = int(MBR.GetWord().data[3]&0b00001111)<<8 + int(MBR.GetWord().data[4])
	return
}

// TakeInstructionFromLeftM takes next instruction from left half of M(x)
func TakeInstructionFromLeftM() {
	PC.SetAddr(MAR.GetAddr())
	FlagIsNextInstructionInIBR = false
}

// TakeInstructionFromRightM takes next instruction from right half of M(x)
func TakeInstructionFromRightM() {
	PC.SetAddr(MAR.GetAddr())
	FlagIsNextInstructionInIBR = false
	FlagLeftInsturctionRequired = false
}
