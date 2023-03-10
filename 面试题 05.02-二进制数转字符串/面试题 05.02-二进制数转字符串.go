package main

import (
	"strings"
)

/*
二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。
*/

// 面试题 05.02-二进制数转字符串
func printBin(num float64) string {
	res := strings.Builder{}
	res.WriteString("0.")
	for res.Len() < 32 && num != 0 {
		num *= 2
		tmp := byte(num)
		res.WriteByte('0' + tmp)
		num -= float64(tmp)
	}
	if res.Len() < 32 {
		return res.String()
	}
	return "ERROR"
}
