package simulator

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

// One word has 40 bit
type Word struct {
	data []byte
}

func NewWord() *Word {
	return &Word{
		data: make([]byte, 5),
	}
}

func NewWordFromHexStr(data string) (*Word, error) {
	// 尝试将字符串转换为十六进制字节切片
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	if len(bytes) != 5 {
		return nil, fmt.Errorf("fail to convert string to word, data length is not 5 bytes")
	}

	return &Word{
		data: bytes,
	}, nil
}

func (w *Word) DeepCopy() *Word {
	newWord := Word{
		data: make([]byte, 5),
	}
	copy(newWord.data, w.data)
	return &newWord
}

func (w *Word) Opposite() *Word {
	opw := w.DeepCopy()
	opw.data[0] ^= 0b10000000
	return opw
}

func (w *Word) Abs() *Word {
	opw := w.DeepCopy()
	opw.data[0] &= 0b01111111
	return opw
}

func (r *Word) Clear() {
	for i := range r.data {
		r.data[i] = 0b00000000
	}
}

func (r *Word) IsEmpty() bool {
	for i := range r.data {
		if r.data[i] != 0b00000000 {
			return false
		}
	}
	return true
}

func (r *Word) IsNegative() bool {
	return r.data[0]&0b10000000 == 0b10000000
}

func (w *Word) IsValid() bool {
	return len(w.data) == 5
}

// 问就是懒得实现各种运算了，毕竟这不是该仿真的重点
func (w *Word) ToInt64() int64 {
	absw := w.Abs()
	absr := int64(absw.data[4]) + int64(absw.data[3])<<8 + int64(absw.data[2])<<16 + int64(absw.data[1])<<24 + int64(absw.data[0])<<32
	if w.IsNegative() {
		return -absr
	}
	return absr
}

func NewWordFromInt64(v int64) *Word {
	absv := v
	if v < 0 {
		absv = -absv
	}

	w := NewWord()
	w.data[4] = byte(absv & 0b11111111)
	w.data[3] = byte(absv >> 8 & 0b11111111)
	w.data[2] = byte(absv >> 16 & 0b11111111)
	w.data[1] = byte(absv >> 24 & 0b11111111)
	w.data[0] = byte(absv >> 32 & 0b11111111)

	if v < 0 {
		w.data[0] |= 0b10000000
	}

	return w
}

func NewWordFromData(data []byte) *Word {
	return &Word{
		data: data,
	}
}

func (w *Word) Add(w2 *Word) *Word {
	ans := w.ToInt64() + w2.ToInt64()
	return NewWordFromInt64(ans)
}

func (w *Word) Sub(w2 *Word) *Word {
	ans := w.ToInt64() - w2.ToInt64()
	return NewWordFromInt64(ans)
}

func (w *Word) Mul(w2 *Word) (higherWord, lowerWord *Word) {
	b1 := big.NewInt(w.ToInt64())
	b2 := big.NewInt(w2.ToInt64())
	var b3 big.Int
	b3.Mul(b1, b2)
	lower5Bytes := make([]byte, 5)
	higher5Bytes := make([]byte, 5)
	bytesAns := b3.Bytes()
	totalBytes := len(bytesAns)
	for i := range b3.Bytes() {
		index := totalBytes - i - 1 // 结果是大端序,从低字节开始倒序遍历
		if i < 5 {
			lower5Bytes[4-i] = bytesAns[index]
		} else if i < 10 {
			higher5Bytes[9-i] = bytesAns[index]
		} else {
			break
		}
	}

	higherWord = NewWordFromData(higher5Bytes)
	lowerWord = NewWordFromData(lower5Bytes)
	// 设置符号位
	if b3.Sign() < 0 {
		higherWord.data[0] |= 0b10000000
	} else {
		higherWord.data[0] &= 0b01111111
	}
	return
}

func (w *Word) Div(w2 *Word) (quotient, remainder *Word) {
	x1 := w.ToInt64()
	x2 := w2.ToInt64()
	quotientInt64 := x1 / x2
	remainderInt64 := x1 % x2
	quotient = NewWordFromInt64(quotientInt64)
	remainder = NewWordFromInt64(remainderInt64)
	return
}

// Left shift one bit
func (w *Word) LSH() *Word {
	x := w.ToInt64() << 1
	return NewWordFromInt64(x)
}

// Right shift one bit
func (w *Word) RSH() *Word {
	x := w.ToInt64() >> 1
	return NewWordFromInt64(x)
}

// Right shift one bit
func (w *Word) ToHexStr() string {
	var r string
	for _, b := range w.data {
		r = fmt.Sprintf("%s%02X", r, b)
	}
	return r
}
