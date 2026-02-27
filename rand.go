package go_util

import (
	"math/rand"
)

const (
	digits       = "0123456789"
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters      = lowercase + uppercase
	alphanumeric = letters + digits
	emoji        = "😀😃😄😁😆😅😂🤣😊😇🙂🙃😉😌😍🥰😘😗😙😚😋😛😝😜🤪🤨🧐🤓😎🥸🤩🥳😏😒😞😔😟😕🙁☹️😣😖😫😩🥺😢😭😤😠😡🤬🤯😳🥵🥶😱😨😰😥😓🤗🤔🤭🤫🤥😶😐😑😬🙄😯😦😧😮😲🥱😴🤤😪😵🤐🥴🤢🤮🤧😷🤒🤕🤑🤠😈👿👹👺🤡💩👻💀☠️👽👾🤖🎃😺😸😹😻😼😽🙀😿😾🧠👂👃👄👅 🧠👁👂👃👄👅"
)

// RandStr 生成包含字母和数字的随机字符串
func RandStr(length int) string {
	if length <= 0 {
		return ""
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(result)
}

// RandStrDigit 生成纯数字的随机字符串
func RandStrDigit(length int) string {
	if length <= 0 {
		return ""
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

// RandStrLetter 生成纯字母的随机字符串
func RandStrLetter(length int) string {
	if length <= 0 {
		return ""
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// RandStrEmoji 生成随机表情符号字符串
func RandStrEmoji(length int) string {
	if length <= 0 {
		return ""
	}

	emojiRunes := []rune(emoji)
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = emojiRunes[rand.Intn(len(emojiRunes))]
	}
	return string(result)
}

// RandInt 生成指定范围内的随机整数 [min, max)
func RandInt(min, max int) int {
	if min >= max {
		return min
	}
	return rand.Intn(max-min) + min
}

// RandFloat 生成指定范围内的随机浮点数 [min, max)
func RandFloat(min, max float64) float64 {
	if min >= max {
		return min
	}
	return rand.Float64()*(max-min) + min
}

// RandChoice 从切片中随机选择一个元素
func RandChoice[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[rand.Intn(len(slice))]
}

// RandShuffle 随机打乱切片
func RandShuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
