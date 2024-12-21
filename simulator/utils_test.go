package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInstructionAndAddrListToHexStrList(t *testing.T) {
	myProgram := []*InstructionAndAddr{
		{OpcodeLoadM, 250},                // 010FA
		{OpcodeStoreM, 251},               // 210FB
		{OpcodeLoadM, 250},                // 010FA
		{OpcodeConditionalJumpMLeft, 141}, // 0F08D
		{OpcodeLoadNegativeM, 250},        // 020FA
		{OpcodeStoreM, 251},               // 210FB
	}
	got := ConvertInstructionAndAddrListToHexStrList(myProgram)
	want := []string{"010FA210FB", "010FA0F08D", "020FA210FB"}
	assert.Equal(t, got, want)
}
