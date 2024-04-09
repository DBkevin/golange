package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Uint64ToString 将字符串转换为String
func Uint64ToString(num uint64) string {

	return strconv.FormatUint(num, 10)
}

// StringToUint64 将字符串转换为 Uint64
func StringToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)

	}
	return i
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
