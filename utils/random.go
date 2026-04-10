package utils

import (
	"math/rand"
)

// GenerateCode 生成 N 位纯数字验证码
func GenerateCode(length int) string {
	if length <= 0 {
		length = 6 // 默认6位
	}

	var codes []rune
	for i := 0; i < length; i++ {
		// 生成 0-9 的数字
		codes = append(codes, rune('0'+rand.Intn(10)))
	}

	return string(codes)
}
