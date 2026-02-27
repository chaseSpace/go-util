package ustr

import (
	"strings"
	"unicode"
)

// CamelToUnderScore 将驼峰命名转换为下划线命名
// 例如: "CamelCase" -> "camel_case", "HTMLParser" -> "html_parser"
func CamelToUnderScore(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		// 如果是大写字母
		if unicode.IsUpper(r) {
			// 添加下划线的条件：
			// 1. 不是在开头
			// 2. 前一个字符不是大写，或者
			// 3. 后面还有小写字母（处理连续大写的情况，如HTMLParser）
			if i > 0 {
				if !unicode.IsUpper(runes[i-1]) ||
					(i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
					result.WriteRune('_')
				}
			}
			// 转换为小写
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// UnderScoreToCamel 将下划线命名转换为驼峰命名
// 例如: "under_score" -> "UnderScore", "html_parser" -> "HtmlParser"
func UnderScoreToCamel(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "_")
	var result strings.Builder

	for _, word := range words {
		if word != "" {
			// 首字母大写，其余小写
			runes := []rune(word)
			if len(runes) > 0 {
				result.WriteRune(unicode.ToUpper(runes[0]))
				for i := 1; i < len(runes); i++ {
					result.WriteRune(unicode.ToLower(runes[i]))
				}
			}
		}
	}

	return result.String()
}

// ToKebabCase 将字符串转换为kebab-case格式
// 例如: "CamelCase" -> "camel-case", "under_score" -> "under-score"
func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	runes := []rune(s)
	lastWasDash := false

	for i, r := range runes {
		switch {
		case unicode.IsUpper(r):
			if i > 0 && !lastWasDash {
				prev := runes[i-1]

				// 只有字母或字母与数字之间才考虑分词
				if (unicode.IsLetter(prev) || unicode.IsDigit(prev)) &&
					(unicode.IsLower(prev) ||
						(i+1 < len(runes) && unicode.IsLower(runes[i+1]))) {
					result.WriteRune('-')
					lastWasDash = true
				}
			}

			result.WriteRune(unicode.ToLower(r))
			lastWasDash = false

		case r == '_' || r == '-':
			if !lastWasDash && result.Len() > 0 {
				result.WriteRune('-')
				lastWasDash = true
			}

		default:
			result.WriteRune(r)
			lastWasDash = false
		}
	}

	out := result.String()
	if len(out) > 0 && out[len(out)-1] == '-' {
		out = out[:len(out)-1]
	}

	return out
}
