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

func TestWordToInt64(t *testing.T) {
	test := []struct {
		w       *Word
		wantInt int64
	}{
		{
			w: &Word{
				data: []byte{0, 0, 0b10111100, 0b01100001, 0b01001110}, // 10111100 01100001 01001110
			},
			wantInt: 12345678,
		},
		{
			w: &Word{
				data: []byte{0b10000000, 0, 0b10111100, 0b01100001, 0b01001110},
			},
			wantInt: -12345678,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// to int64
			gotInt := tt.w.ToInt64()
			assert.Equal(t, gotInt, tt.wantInt)
			// from int64
			gotw := NewWordFromInt64(tt.wantInt)
			assert.Equal(t, gotw, tt.w)
		})
	}
}
