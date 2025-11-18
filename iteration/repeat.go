package iteration

import (
	"strings"
)

const repeatCount = 5

func Repeat(character string, times int) string {
	var repeated strings.Builder // var就是 explicit declaration。
	for i := 0; i < times; i++ {
		//go之中的字符串是不可变的，这就意味着连接操作，会消耗大量内存
		//repeated += character
		repeated.WriteString(character)
	}
	return repeated.String()
}
