package ustr

func TruncateUTF8(s string, length int, appendTail ...string) string {
	if length == 0 {
		return ""
	}

	runes := []rune(s)
	if len(runes) > length {
		if len(appendTail) > 0 {
			return string(runes[:length]) + appendTail[0]
		}
		return string(runes[:length]) // 截断到指定长度
	}

	return s
}
