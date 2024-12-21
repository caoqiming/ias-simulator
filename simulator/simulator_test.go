package simulator

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSetInstructions(t *testing.T) {
	test := []struct {
		data      []*InstructionAndAddr
		startAddr int
		want      []*Word
	}{
		{
			data: []*InstructionAndAddr{
				{
					OpcodeLoadMQ,
					100,
				},
				{
					OpcodeConditionalJumpMLeft,
					4000,
				},
			},
			startAddr: 0,
			want: []*Word{
				{
					[]byte{0b00001010, 0b00000110, 0b01000000, 0b11111111, 0b10100000},
				},
			},
		},
		{
			data: []*InstructionAndAddr{
				{
					OpcodeLoadMQ,
					100,
				},
				{
					OpcodeConditionalJumpMLeft,
					4000,
				},
				{
					OpcodeLSH,
					2025,
				},
			},
			startAddr: 0,
			want: []*Word{
				{
					[]byte{0b00001010, 0b00000110, 0b01000000, 0b11111111, 0b10100000},
				},
				{
					[]byte{0b00010100, 0b01111110, 0b10010000, 0, 0},
				},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i),
			func(t *testing.T) {
				SetInstructions(tt.data, tt.startAddr)
				for i := range tt.data {
					if i%2 != 0 {
						continue
					}
					if i+1 < len(tt.data) {
						// 一次检查两个指令
						got := DirectRead(i + tt.startAddr)
						assert.Equal(t, got, tt.want[i])
					}
				}
				if len(tt.data)%2 != 0 {
					// 检查最后一个指令
					got := DirectRead(tt.startAddr + len(tt.data)/2)
					assert.Equal(t, got, tt.want[len(tt.data)/2])
				}
			})
	}
}

func TestTakeNextInstruction(t *testing.T) {
	test := []struct {
		data      []*InstructionAndAddr
		startAddr int
		want      []*Word
	}{
		{
			data: []*InstructionAndAddr{
				{
					OpcodeLoadMQ,
					100,
				},
				{
					OpcodeConditionalJumpMLeft,
					4000,
				},
			},
			startAddr: 0,
		},
		{
			data: []*InstructionAndAddr{
				{
					OpcodeLoadMQ,
					123,
				},
				{
					OpcodeConditionalJumpMLeft,
					4090,
				},
				{
					OpcodeJumpMRight,
					1231,
				},
				{
					OpcodeSubM,
					1234,
				},
				{
					OpcodeDivideM,
					1312,
				},
				{
					OpcodeAddAbsM,
					1312,
				},
				{
					OpcodeAddAbsM,
					1312,
				},
			},
			startAddr: 2025,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i),
			func(t *testing.T) {
				SetInstructions(tt.data, tt.startAddr)
				PC.SetAddr(tt.startAddr)
				for i := range tt.data {
					TakeNextInstruction()
					assert.Equal(t, IR.Read(), tt.data[i].OpCode)
					assert.Equal(t, MAR.GetAddr(), tt.data[i].Addr)
				}
			})
	}
}
