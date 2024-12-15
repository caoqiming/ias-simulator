package simulator

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestInstructionLoadMQ(t *testing.T) {
	instruction := &InstructionLoadMQ{}
	MQ.data.data[0] = 0b00001111
	instruction.Run()
	assert.Equal(t, AC.data.data[0], MQ.data.data[0])
}

func TestInstructionLoadMToMQ(t *testing.T) {
	instruction := &InstructionLoadMToMQ{}
	MAR.SetAddr(1000)
	MBR.SetWord(&Word{data: []byte{0b00001111, 0, 0, 0, 0}})
	memory.Write()
	instruction.Run()
	assert.Equal(t, MBR.data.data[0], byte(0b00001111))
	assert.Equal(t, MQ.data.data[0], byte(0b00001111))
}

func TestInstructionStoreM(t *testing.T) {
	instruction := &InstructionStoreM{}
	MAR.SetAddr(1024)
	AC.SetWord(&Word{data: []byte{0b11111111, 0, 0, 0, 0}})
	instruction.Run()
	assert.Equal(t, MBR.data.data[0], byte(0b11111111))
	assert.Equal(t, memory.DirectRead(1024).data[0], byte(0b11111111))
}

func TestInstructionLoadM(t *testing.T) {
	instruction := &InstructionLoadM{}
	MAR.SetAddr(2025)
	AC.SetWord(&Word{data: []byte{0b11111111, 0, 0, 0, 0}})
	memory.DirectWrite(2025, &Word{data: []byte{0b01101001, 0, 0, 0, 0}})
	instruction.Run()
	assert.Equal(t, MBR.data.data[0], byte(0b01101001))
	assert.Equal(t, AC.data.data[0], byte(0b01101001))
}

func TestGetInstructionFromMBR(t *testing.T) {
	test := []struct {
		mbr           *Word
		wantLeftCode  byte
		wantLeftAddr  int
		wantRightCode byte
		wantRightAddr int
	}{
		{
			mbr:           &Word{data: []byte{0b00100001, 0b00000001, 0b00100001, 0b00100001, 0b00100001}},
			wantLeftCode:  0b00100001,
			wantLeftAddr:  0b000000010010,
			wantRightCode: 0b00010010,
			wantRightAddr: 0b000100100001,
		},
		{
			mbr:           &Word{data: []byte{0b00010010, 0b10010110, 0b11110000, 0b11000010, 0b01000001}},
			wantLeftCode:  0b00010010,
			wantLeftAddr:  0b100101101111,
			wantRightCode: 0b00001100,
			wantRightAddr: 0b001001000001,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			MBR.SetWord(tt.mbr)
			gotLeftCode, gotLeftAddr := GetLeftInstructionFromMBR()
			assert.Equal(t, gotLeftCode, tt.wantLeftCode)
			assert.Equal(t, gotLeftAddr, tt.wantLeftAddr)
			gotRightCode, gotRightAddr := GetRightInstructionFromMBR()
			assert.Equal(t, gotRightCode, tt.wantRightCode)
			assert.Equal(t, gotRightAddr, tt.wantRightAddr)
		})
	}

}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}
