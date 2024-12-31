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

// print 当前各个寄存器的数据，用于debug
func PrintStatus() {
	fmt.Println("=====ias-simulator-status=====")
	fmt.Printf("PC: %d IR: %d\n", PC.GetAddr(), IR.Read())
	fmt.Printf("MAR: %d MBR: %d\n", MAR.GetAddr(), MBR.GetWord().ToInt64())
	fmt.Printf("AC: %d MQ: %d \n", AC.GetWord().ToInt64(), MQ.GetWord().ToInt64())
	fmt.Println("=============================")
}

// 返回各个寄存器的数据
func SPrintStatus() string {
	var result string
	result += fmt.Sprintf("PC:  %10d  IR: %10d\n", PC.GetAddr(), IR.Read())
	result += fmt.Sprintf("MAR: %10d MBR: %10d\n", MAR.GetAddr(), MBR.GetWord().ToInt64())
	result += fmt.Sprintf("AC:  %10d  MQ: %10d\n", AC.GetWord().ToInt64(), MQ.GetWord().ToInt64())
	return result
}

// 将程序转化为16进制的格式
func ConvertInstructionAndAddrListToHexStrList(data []*InstructionAndAddr) []string {
	result := []string{}
	for i := range data {
		if i%2 != 0 {
			continue
		}

		w := NewWord()
		// 写入当前指令到左边
		w.data[0] = data[i].OpCode
		w.data[1] = byte(data[i].Addr >> 4)
		w.data[2] = byte(data[i].Addr&0b00001111) << 4
		if i+1 < len(data) {
			// 写入下个指令到右边
			higherByte, lowerByte := ConvertIntToTwoByte(data[i+1].Addr)
			w.data[2] += data[i+1].OpCode & 0b11110000 >> 4
			w.data[3] = data[i+1].OpCode & 0b00001111 << 4
			w.data[3] += higherByte
			w.data[4] = lowerByte
		}
		result = append(result, w.ToHexStr())
	}

	return result
}
