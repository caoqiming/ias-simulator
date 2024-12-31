package example

import (
	"fmt"
	"testing"

	s "github.com/caoqiming/ias-simulator/simulator"
	"github.com/stretchr/testify/assert"
)

// 这段程序来自 Computer Organization and Architecture Designing for Performance, Eleventh Edition Execrise 1.1 a
func TestExecrise1_1_a(t *testing.T) {
	// 这一段程序是求和1到N的等差数列，假设结果用一个word就可以表示
	// 不妨假设数字N输入在地址 100，我们将结果输出到地址101
	// 这里直接用公式 N*（N+1）/2
	// 初始化
	// 将 1 写到 102
	myProgram := []*s.InstructionAndAddr{
		{s.OpcodeLoadM, 100},     // N -> AC
		{s.OpcodeAddM, 102},      // AC = AC + 1
		{s.OpcodeStoreM, 102},    // 存 N+1 到 102
		{s.OpcodeLoadMToMQ, 102}, // 读 N+1 到 MQ
		{s.OpcodeMultiplyM, 100}, // MQ * N
		{s.OpcodeLoadMQ, 0},      // MQ -> AC
		{s.OpcodeRSH, 0},         // AC = AC / 2
		{s.OpcodeStoreM, 101},
	}
	test := []struct {
		ori  int64
		want int64
	}{
		{
			10,
			55,
		},
		{
			100,
			5050,
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
				s.DirectWrite(102, s.NewWordFromInt64(1))
				s.DirectWrite(100, s.NewWordFromInt64(tt.ori))
				s.SetInstructions(myProgram, 0)
				s.PC.SetAddr(0)
				err := s.Execute(s.WithMaxSteps(len(myProgram)))
				if err != nil {
					s.PrintStatus()
				}
				assert.Nil(t, err)
				got := s.DirectRead(101)
				assert.Equal(t, tt.want, got.ToInt64())
			})
	}
}

// 这段程序来自 Computer Organization and Architecture Designing for Performance, Eleventh Edition, Execrise 1.1 b
func TestExecrise1_1_b(t *testing.T) {
	// 这一段程序是求和1到N的等差数列，假设结果用一个word就可以表示
	// 不妨假设数字N输入在地址 100，我们将结果输出到地址101
	// 这里老老实实挨个加
	// 初始化
	// 将 0 写到 102（用于迭代）, 将 1 写到 103（用于常量1）, 将 0 写到 101
	myProgram := []*s.InstructionAndAddr{
		// 判断 120 的值是否达到 N
		{s.OpcodeLoadM, 102},              // at 0
		{s.OpcodeSubM, 100},               // 与 N 相减
		{s.OpcodeConditionalJumpMLeft, 5}, // 102 已经达到了 N，结束程序, at 1
		// 102 的值加一
		{s.OpcodeLoadM, 102},
		{s.OpcodeAddM, 103}, // at 2
		{s.OpcodeStoreM, 102},
		// 计算一次加法
		{s.OpcodeAddM, 101},   // 计算新的 sum, at 3
		{s.OpcodeStoreM, 101}, // 储存新的 sum, at
		// 循环
		{s.OpcodeJumpMLeft, 0}, // at 4
	}
	test := []struct {
		ori  int64
		want int64
	}{
		{
			10,
			55,
		},
		{
			100,
			5050,
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
				s.DirectWrite(101, s.NewWordFromInt64(0))
				s.DirectWrite(102, s.NewWordFromInt64(0))
				s.DirectWrite(103, s.NewWordFromInt64(1))
				s.DirectWrite(100, s.NewWordFromInt64(tt.ori))
				s.SetInstructions(myProgram, 0)
				s.PC.SetAddr(0)
				err := s.Execute(s.WithHaultAt(5))
				if err != nil {
					s.PrintStatus()
				}
				assert.Nil(t, err)
				got := s.DirectRead(101)
				assert.Equal(t, tt.want, got.ToInt64())
			})
	}
}
