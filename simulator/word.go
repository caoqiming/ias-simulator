package simulator

// One word has 40 bit
type Word struct {
	data []byte
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
		return w.Opposite()
	}

	return w
}

func (w *Word) Add(w2 *Word) *Word {
	ans := w.ToInt64() + w2.ToInt64()
	return NewWordFromInt64(ans)
}

func (w *Word) Sub(w2 *Word) *Word {
	ans := w.ToInt64() - w2.ToInt64()
	return NewWordFromInt64(ans)
}

func NewWord() *Word {
	return &Word{
		data: make([]byte, 5),
	}
}
