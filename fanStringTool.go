package fanTool

import (
	"strconv"
	"unicode/utf8"
)

// IsNumeric 判断是否为数字
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsInt 判断是否为整数
func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// IsFloat 判断是否为浮点数
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// StringToInt 字符串转int
func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// StringToInt8 字符串转int8
func StringToInt8(s string) int8 {
	num, _ := strconv.ParseInt(s, 10, 8)
	return int8(num)
}

// StringToInt32 字符串转int16
func StringToInt32(s string) int32 {
	num, _ := strconv.ParseInt(s, 10, 32)
	return int32(num)
}

// StringToInt64 字符串转int64
func StringToInt64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

// StringToFloat32 字符串转float32
func StringToFloat32(s string) float32 {
	num, _ := strconv.ParseFloat(s, 32)
	return float32(num)
}

// StringToFloat64 字符串转float64
func StringToFloat64(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

// StringToBool 字符串转bool
func StringToBool(s string) bool {
	if s == "true" || s == "1" {
		return true
	}
	return false
}

// StringCount 字符串长度
func StringCount(s string) int {
	return utf8.RuneCountInString(s)
}
