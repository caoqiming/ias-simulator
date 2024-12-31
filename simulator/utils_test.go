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

	myProgram = []*InstructionAndAddr{
		{OpcodeLoadM, 100},     // N -> AC
		{OpcodeAddM, 102},      // AC = AC + 1
		{OpcodeStoreM, 102},    // 存 N+1 到 102
		{OpcodeLoadMToMQ, 102}, // 读 N+1 到 MQ
		{OpcodeMultiplyM, 100}, // MQ * N
		{OpcodeLoadMQ, 0},      // MQ -> AC
		{OpcodeRSH, 0},         // AC = AC / 2
		{OpcodeStoreM, 101},
	}
	got = ConvertInstructionAndAddrListToHexStrList(myProgram)
	want = []string{"0106405066", "2106609066", "0B0640A000", "1500021065"}
	assert.Equal(t, got, want)
}
