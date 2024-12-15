package simulator

import "fmt"

func ConvertIntToTwoByte(addr int) (higherByte, lowerByte byte) {
	if addr < 0 || addr >= 4096 {
		panic(fmt.Sprintf("invalid addr %d", addr))
	}
	higherByte = byte(addr >> 8)
	lowerByte = byte(addr & 0b11111111)
	return
}

func ConvertTwoByteToInt(higherByte, lowerByte byte) int {
	return int(higherByte&0b00001111)<<8 + int(lowerByte)
}
