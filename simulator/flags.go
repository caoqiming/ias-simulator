package simulator

var (
	FlagIsNextInstructionInIBR  bool
	FlagLeftInsturctionRequired bool // 下一个指令应该执行左边的指令吗
)

func initFlag() {
	FlagIsNextInstructionInIBR = false
	FlagLeftInsturctionRequired = true
}
