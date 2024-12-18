package simulator

import "fmt"

var (
	memory *Memory
)

type Memory struct {
	words []Word
}

// Read reads the memory to MBR
func (m *Memory) Read() {
	addr := MAR.GetAddr()
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %v", addr))
	}
	MBR.SetWord(m.words[addr].DeepCopy())
}

// Write writes MBR to memory
func (m *Memory) Write() {
	addr := MAR.GetAddr()
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %d", addr))
	}

	w := MBR.GetWord()
	if !w.IsValid() {
		panic(fmt.Sprintf("word try to write at %d is not valid", addr))
	}

	m.words[addr] = *w.DeepCopy()
}

// Write writes MBR[28:39] to M(X)[8:19]
func (m *Memory) WriteLeftAddr() {
	addr := MAR.GetAddr()
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %d", addr))
	}

	w := MBR.GetWord()
	if !w.IsValid() {
		panic(fmt.Sprintf("word try to write at %d is not valid", addr))
	}

	// 8~15
	m.words[addr].data[1] = (MBR.GetWord().data[3]&0b00001111)<<4 + (MBR.GetWord().data[4]&0b11110000)>>4
	// 16~19
	m.words[addr].data[2] = m.words[addr].data[2]&0b00001111 + (MBR.GetWord().data[4]&0b00001111)<<4
}

// Write writes MBR[28:39] to M(X)[28:39]
func (m *Memory) WriteRightAddr() {
	addr := MAR.GetAddr()
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %d", addr))
	}

	w := MBR.GetWord()
	if !w.IsValid() {
		panic(fmt.Sprintf("word try to write at %d is not valid", addr))
	}

	// 28~31
	m.words[addr].data[3] = m.words[addr].data[3]&0b11110000 + MBR.GetWord().data[3]&0b00001111
	// 32~39
	m.words[addr].data[4] = MBR.GetWord().data[4]
}

// 直接读取内存，只应该用于测试或初始化
func (m *Memory) DirectRead(addr int) *Word {
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %v", addr))
	}
	return m.words[addr].DeepCopy()
}

// 直接写入内存，只应该用于测试或初始化
func (m *Memory) DirectWrite(addr int, w *Word) {
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %v", addr))
	}
	m.words[addr] = *w.DeepCopy()
}

func initMemory() {
	memory = &Memory{}
	// The address is 12-bit, so a maximum of 4096 memory addresses are supported
	memory.words = make([]Word, 4096)
	// Each word has 40 bit, that is 5 byte
	for _, w := range memory.words {
		w.data = make([]byte, 5)
	}
}
