package simulator

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMemory(t *testing.T) {
	valueFrom := &Word{
		data: []byte{0, 0, 0, 0b00001111, 0},
	}
	want := &Word{
		data: []byte{0, 0, 0, 0b00001111, 0},
	}
	MAR.SetAddr(10)
	MBR.SetWord(valueFrom)
	memory.Write()
	memory.Read()
	got := MBR.GetWord()
	assert.Equal(t, got, want)
	// memory should not change when value from has changed
	valueFrom.data[0] = 0b10101010
	memory.Read()
	got2 := MBR.GetWord()
	assert.Equal(t, got2, want)
}

func TestWordOpesite(t *testing.T) {
	test := []struct {
		oriWprd *Word
		want    *Word
	}{
		{
			oriWprd: &Word{
				data: []byte{0b01010101, 0, 0, 0, 0},
			},
			want: &Word{
				data: []byte{0b11010101, 0, 0, 0, 0},
			},
		},
		{
			oriWprd: &Word{
				data: []byte{0b11010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
			want: &Word{
				data: []byte{0b01010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.oriWprd.Opposite(), tt.want)
		})
	}
}

func TestWordAbs(t *testing.T) {
	test := []struct {
		oriWprd *Word
		want    *Word
	}{
		{
			oriWprd: &Word{
				data: []byte{0b01010101, 0, 0, 0, 0},
			},
			want: &Word{
				data: []byte{0b01010101, 0, 0, 0, 0},
			},
		},
		{
			oriWprd: &Word{
				data: []byte{0b1010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
			want: &Word{
				data: []byte{0b01010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, tt.oriWprd.Abs(), tt.want)
		})
	}
}

func TestWordClear(t *testing.T) {
	test := []struct {
		oriWprd *Word
		want    *Word
	}{
		{
			oriWprd: &Word{
				data: []byte{0b01010101, 0, 0, 0, 0},
			},
			want: &Word{
				data: []byte{0, 0, 0, 0, 0},
			},
		},
		{
			oriWprd: &Word{
				data: []byte{0b11010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
			want: &Word{
				data: []byte{0, 0, 0, 0, 0},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tt.oriWprd.Clear()
			assert.Equal(t, tt.oriWprd, tt.want)
		})
	}
}

func TestWordIsNegative(t *testing.T) {
	test := []struct {
		w    *Word
		want bool
	}{
		{
			w: &Word{
				data: []byte{0b01010101, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			w: &Word{
				data: []byte{0b11010101, 0b11111111, 0b01011001, 0b00000000, 0b10110010},
			},
			want: true,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := tt.w.IsNegative()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestMemoryWriteLeftAddr(t *testing.T) {
	test := []struct {
		addr int
		oriW *Word
		MBR  *Word
		want *Word
	}{
		{
			addr: 25,
			oriW: &Word{
				data: []byte{0, 0, 0, 0, 0},
			},
			MBR: &Word{
				data: []byte{0, 0, 0, 0b00001111, 0b11111111},
			},
			want: &Word{
				data: []byte{0, 0b11111111, 0b11110000, 0, 0},
			},
		},
		{
			addr: 2025,
			oriW: &Word{
				data: []byte{0b00111001, 0b00101111, 0b01100111, 0b00001111, 0b11010011},
			},
			MBR: &Word{
				data: []byte{0b01101001, 0b01101010, 0b11010110, 0b10101010, 0b00101110},
			},
			want: &Word{
				data: []byte{0b00111001, 0b10100010, 0b11100111, 0b00001111, 0b11010011},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			memory.DirectWrite(tt.addr, tt.oriW)
			MAR.SetAddr(tt.addr)
			MBR.SetWord(tt.MBR)
			memory.WriteLeftAddr()
			got := memory.DirectRead(tt.addr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestMemoryWriteRightAddr(t *testing.T) {
	test := []struct {
		addr int
		oriW *Word
		MBR  *Word
		want *Word
	}{
		{
			addr: 25,
			oriW: &Word{
				data: []byte{0, 0, 0, 0, 0},
			},
			MBR: &Word{
				data: []byte{0, 0, 0, 0b00001111, 0b11111111},
			},
			want: &Word{
				[]byte{0, 0, 0, 0b00001111, 0b11111111},
			},
		},
		{
			addr: 2025,
			oriW: &Word{
				data: []byte{0b00111001, 0b00101111, 0b01100111, 0b00001111, 0b11010011},
			},
			MBR: &Word{
				data: []byte{0b01101001, 0b01101010, 0b11010110, 0b10101010, 0b00101110},
			},
			want: &Word{
				data: []byte{0b00111001, 0b00101111, 0b01100111, 0b00001010, 0b00101110},
			},
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			memory.DirectWrite(tt.addr, tt.oriW)
			MAR.SetAddr(tt.addr)
			MBR.SetWord(tt.MBR)
			memory.WriteRightAddr()
			got := memory.DirectRead(tt.addr)
			assert.Equal(t, got, tt.want)
		})
	}
}
