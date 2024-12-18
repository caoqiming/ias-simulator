package simulator

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

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
		{
			w: &Word{
				data: []byte{0b11000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000},
			},
			wantInt: -274877906944,
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

func TestWordMul(t *testing.T) {
	test := []struct {
		x1       int64
		x2       int64
		wantHigh *Word
		wantLow  *Word
	}{
		{
			x1:       274888999555,
			x2:       274333555777,
			wantHigh: NewWordFromData([]byte{0b00001111, 0b11111000, 0b00001101, 0b10110010, 0b00111101}),
			wantLow:  NewWordFromData([]byte{0b00111000, 0b01011111, 0b01001001, 0b01110111, 0b01000011}),
		},
		{
			x1:       -274877906944,
			x2:       274877906944,
			wantHigh: NewWordFromData([]byte{0b10010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}),
			wantLow:  NewWordFromData([]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}),
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// to int64
			w1 := NewWordFromInt64(tt.x1)
			w2 := NewWordFromInt64(tt.x2)
			gotHigh, gotLow := w1.Mul(w2)
			// b1 := big.NewInt(tt.x1)
			// b2 := big.NewInt(tt.x2)
			// b1.Mul(b1, b2)
			// fmt.Printf("%08b\n", b1.Bytes())
			assert.Equal(t, gotHigh, tt.wantHigh)
			assert.Equal(t, gotLow, tt.wantLow)
		})
	}
}

func TestWordDiv(t *testing.T) {
	test := []struct {
		x1           int64
		x2           int64
		wantQuotient int64
		wantRemain   int64
	}{
		{
			x1:           10,
			x2:           3,
			wantQuotient: 3,
			wantRemain:   1,
		},
		{
			x1:           1000,
			x2:           7,
			wantQuotient: 142,
			wantRemain:   6,
		},
		{
			x1:           12345678,
			x2:           -9,
			wantQuotient: -1371742,
			wantRemain:   0,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// to int64
			w1 := NewWordFromInt64(tt.x1)
			w2 := NewWordFromInt64(tt.x2)
			gotQuotientW, gotRemainderW := w1.Div(w2)
			assert.Equal(t, gotQuotientW.ToInt64(), tt.wantQuotient)
			assert.Equal(t, gotRemainderW.ToInt64(), tt.wantRemain)
		})
	}
}

func TestWordLSH(t *testing.T) {
	test := []struct {
		x    int64
		want int64
	}{
		{
			x:    10,
			want: 20,
		},
		{
			x:    33333,
			want: 66666,
		},
		{
			x:    -12391920,
			want: -24783840,
		},
		{
			x:    0,
			want: 0,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// to int64
			w := NewWordFromInt64(tt.x)
			gotW := w.LSH()
			assert.Equal(t, gotW.ToInt64(), tt.want)
		})
	}
}

func TestWordRSH(t *testing.T) {
	test := []struct {
		x    int64
		want int64
	}{
		{
			x:    20,
			want: 10,
		},
		{
			x:    66666,
			want: 33333,
		},
		{
			x:    -24783840,
			want: -12391920,
		},
		{
			x:    0,
			want: 0,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// to int64
			w := NewWordFromInt64(tt.x)
			gotW := w.RSH()
			assert.Equal(t, gotW.ToInt64(), tt.want)
		})
	}
}
