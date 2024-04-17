package fanTool

import (
	"strconv"
)

// intToString 将数字转换为字符串
func intToString(num interface{}) string {
	switch v := num.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(num.(int64), 10)
	case int16:
		return strconv.FormatInt(num.(int64), 10)
	case int32:
		return strconv.FormatInt(num.(int64), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return num.(string)
	default:
		return ""
	}
}
