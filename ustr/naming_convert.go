package ustr

import (
	"strings"
	"unicode"
)

// ToSnakeCase 将字符串转换为snake_case格式（支持多种输入格式）
// 例如: "CamelCase" -> "camel_case", "html-parser" -> "html_parser", "already_snake" -> "already_snake"
func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// 统一分隔符处理
	var normalized strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		switch {
		case unicode.IsUpper(r):
			// 大写字母前添加下划线（除非是第一个字符或前一个是下划线/特殊字符）
			if i > 0 {
				prev := runes[i-1]
				// 只在字母或数字之间添加下划线
				if (unicode.IsLetter(prev) || unicode.IsDigit(prev)) &&
					prev != '_' &&
					(unicode.IsLower(prev) || (i+1 < len(runes) && unicode.IsLower(runes[i+1]))) {
					normalized.WriteRune('_')
				}
			}
			normalized.WriteRune(unicode.ToLower(r))
		case r == '-':
			// 连字符转下划线
			normalized.WriteRune('_')
		default:
			normalized.WriteRune(r)
		}
	}

	// 处理连续下划线
	result := normalized.String()
	var final strings.Builder
	prevWasUnderscore := false

	for _, r := range result {
		if r == '_' {
			if !prevWasUnderscore {
				final.WriteRune(r)
				prevWasUnderscore = true
			}
		} else {
			final.WriteRune(r)
			prevWasUnderscore = false
		}
	}

	return strings.Trim(final.String(), "_")
}

// ToCamelCase 将字符串转换为驼峰命名格式（支持多种输入格式）
// 例如: "under_score" -> "UnderScore", "html-parser" -> "HtmlParser", "camelCase" -> "CamelCase"
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	// 统一分隔符为下划线
	s = strings.ReplaceAll(s, "-", "_")

	// 处理驼峰转下划线的情况
	var normalized strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 {
			// 在大写字母前添加下划线（除非前一个也是大写）
			if !unicode.IsUpper(runes[i-1]) {
				normalized.WriteRune('_')
			}
		}
		normalized.WriteRune(unicode.ToLower(r))
	}

	// 按下划线分割并转换为驼峰
	words := strings.Split(normalized.String(), "_")
	var result strings.Builder

	for _, word := range words {
		if word != "" {
			// 首字母大写
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
