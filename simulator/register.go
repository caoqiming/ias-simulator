package simulator

var (
	// Accumulator register, 40 bit
	AC *Register
	// Multiply-quotient register, 40 bit
	MQ *Register
	// Memory buffer register, 40 bit
	MBR *Register
	// Memory address register, 12 bit
	MAR *AddressRegister
	// Program counter, 12 bit
	PC *AddressRegister
	// Instruction buffer register, 20bit
	IBR *InstructionBufferRegister
	// Instruction register, 8 bit
	IR *InstructionRegister
)

type Register struct {
	data *Word
}

func (r *Register) GetWord() *Word {
	return r.data.DeepCopy()
}

func (r *Register) SetWord(w *Word) {
	r.data = w
}

func (r *Register) Clear() {
	r.data.Clear()
}

func (r *Register) IsEmpty() bool {
	return r.data.IsEmpty()
}

func (r *Register) IsNegative() bool {
	return r.data.IsNegative()
}

func NewRegister() *Register {
	return &Register{
		data: NewWord(),
	}
}

type InstructionBufferRegister struct {
	data []byte // 用三个字节表示 20 bit，data[0]表示code，其余表示地址，data[1][4:7] 作为高位，data[2][0,8]作为低位 (BigEndian), data[1][0:3]不使用
}

func (r *InstructionBufferRegister) IsEmpty() bool {
	for i := range r.data {
		if r.data[i] != 0b00000000 {
			return false
		}
	}
	return true
}

func (r *InstructionBufferRegister) Clear() {
	for i := range r.data {
		r.data[i] = 0b00000000
	}
}

func (r *InstructionBufferRegister) Write(code byte, addr int) {
	r.data[0] = code
	higherByte, lowerByte := ConvertIntToTwoByte(addr)
	r.data[1] = byte(higherByte)
	r.data[2] = byte(lowerByte)
}

func (r *InstructionBufferRegister) Read() (code byte, addr int) {
	code = r.data[0]
	addr = ConvertTwoByteToInt(r.data[1], r.data[2])
	return
}

func NewInstructionBufferRegister() *InstructionBufferRegister {
	return &InstructionBufferRegister{
		data: make([]byte, 3),
	}
}

type AddressRegister struct {
	data []byte // 用两个字节表示 12 bit，取data[0][4:7] 作为高位，data[1][0,8]作为低位 (BigEndian), data[0][0:3]不使用
}

func (r *AddressRegister) GetAddr() int {
	return ConvertTwoByteToInt(r.data[0], r.data[1])
}

func (r *AddressRegister) SetAddr(addr int) {
	higherByte, lowerByte := ConvertIntToTwoByte(addr)
	r.data[0] = byte(higherByte)
	r.data[1] = byte(lowerByte)
}

func (r *AddressRegister) Increase() {
	r.SetAddr(r.GetAddr() + 1)
}

func NewAddressRegister() *AddressRegister {
	return &AddressRegister{
		data: make([]byte, 2),
	}
}

type InstructionRegister struct {
	code byte // 用一个字节表示 opcode
}

func (r *InstructionRegister) Read() byte {
	return r.code
}

func (r *InstructionRegister) Write(code byte) {
	r.code = code
}

func NewInstructionRegister() *InstructionRegister {
	return &InstructionRegister{
		code: 0b00000000,
	}
}

func initRegister() {
	AC = NewRegister()
	MQ = NewRegister()
	MBR = NewRegister()

	IBR = NewInstructionBufferRegister()
	IR = NewInstructionRegister()

	MAR = NewAddressRegister()
	PC = NewAddressRegister()
}
