package example

import (
	"fmt"
	"testing"

	s "github.com/caoqiming/ias-simulator/simulator"
	"github.com/stretchr/testify/assert"
)

// 这段程序来自 Computer Organization and Architecture Designing for Performance, Eleventh Edition Execrise 1.4
func TestExecrise1_4(t *testing.T) {
	// 这一段程序是将地址250的值的绝对值放到了地址251
	myProgram := []*s.InstructionAndAddr{
		{s.OpcodeLoadM, 250},                // 010FA at 138
		{s.OpcodeStoreM, 251},               // 210FB
		{s.OpcodeLoadM, 250},                // 010FA at 139
		{s.OpcodeConditionalJumpMLeft, 141}, // 0F08D
		{s.OpcodeLoadNegativeM, 250},        // 020FA at 140
		{s.OpcodeStoreM, 251},               // 210FB
	}
	test := []struct {
		ori  int64
		want int64
	}{
		{
			10,
			10,
		},
		{
			-19,
			19,
		},
		{
			0,
			0,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i),
			func(t *testing.T) {
				s.Init()
				s.DirectWrite(250, s.NewWordFromInt64(tt.ori))
				s.SetInstructions(myProgram, 138)
				s.PC.SetAddr(138)
				err := s.Execute(s.WithMaxSteps(6), s.WithHaultAt(141))
				if err != nil {
					s.PrintStatus()
				}
				assert.Nil(t, err)
				got := s.DirectRead(251)
				assert.Equal(t, tt.want, got.ToInt64())
			})
	}
}

func TestMain(m *testing.M) {
	s.Init()
	m.Run()
}
