package cli

type ProgramConfig struct {
	ProgramPath           string          `yaml:"programPath"`
	ProgramInHexFormat    []string        `yaml:"programInHexFormat"`
	ProgramAddr           int             `yaml:"programAddr"`           // 程序从内存的这个地址开始写入
	ProgramCounterStartAt int             `yaml:"programCounterStartAt"` // 程序从内存的这个地址开始运行
	MaxSteps              int             `yaml:"maxSteps"`              // 程序运行指令个数的上限
	HaltAt                int             `yaml:"haltAt"`                // 程序运行到这个地址就会停止（该地址的指令不会被执行）
	MemorySettings        []MemorySetting `yaml:"memorySettings"`        // 初始化内存数据
}

type MemorySetting struct {
	Addr    int      `yaml:"addr"`
	Content []string `yaml:"content"`
}
