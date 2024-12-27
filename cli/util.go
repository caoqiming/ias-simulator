package cli

import (
	"fmt"
	"regexp"
	"strings"
)

// 将输入根据换行符切分，校验每一个子字符串是否是16进制表示的40个bit
func splitAndValidateProgram(input string) ([]string, error) {
	// 根据换行符分割字符串
	lines := strings.Split(input, "\n")
	var result []string
	hexPattern := regexp.MustCompile(`^[0-9A-Fa-f]{10}$`)
	for _, line := range lines {
		// 去除空格
		trimmed := strings.ReplaceAll(line, " ", "")
		if trimmed == "" {
			continue
		}
		// 检查是否是长度为 10 的 16 进制数
		if hexPattern.MatchString(trimmed) {
			result = append(result, trimmed)
		} else {
			return nil, fmt.Errorf("invalid data: %s", trimmed)
		}
	}
	return result, nil
}
