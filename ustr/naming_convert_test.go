package ustr

import "testing"

func TestToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "驼峰转下划线",
			input:    "CamelCase",
			expected: "camel_case",
		},
		{
			name:     "连续大写",
			input:    "HTMLParser",
			expected: "html_parser",
		},
		{
			name:     "连字符转下划线",
			input:    "html-parser",
			expected: "html_parser",
		},
		{
			name:     "已有的下划线",
			input:    "already_snake_case",
			expected: "already_snake_case",
		},
		{
			name:     "混合格式",
			input:    "Camel-Case_HTML",
			expected: "camel_case_html",
		},
		{
			name:     "数字混合",
			input:    "Version2Update",
			expected: "version2_update",
		},
		{
			name:     "全大写",
			input:    "UPPERCASE",
			expected: "uppercase",
		},
		{
			name:     "全小写",
			input:    "lowercase",
			expected: "lowercase",
		},
		{
			name:     "空字符串",
			input:    "",
			expected: "",
		},
		{
			name:     "单个字符",
			input:    "A",
			expected: "a",
		},
		{
			name:     "连续下划线",
			input:    "multi__underscore___test",
			expected: "multi_underscore_test",
		},
		{
			name:     "前导后导下划线",
			input:    "_leading_trailing_",
			expected: "leading_trailing",
		},
		{
			name:     "特殊字符保持",
			input:    "Test@Case#123",
			expected: "test@case#123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "下划线转驼峰",
			input:    "under_score",
			expected: "UnderScore",
		},
		{
			name:     "连字符转驼峰",
			input:    "html-parser",
			expected: "HtmlParser",
		},
		{
			name:     "小驼峰转大驼峰",
			input:    "camelCase",
			expected: "CamelCase",
		},
		{
			name:     "_already_camel",
			input:    "_already_camel",
			expected: "AlreadyCamel",
		},
		{
			name:     "mixed-format_string",
			input:    "mixed-format_string",
			expected: "MixedFormatString",
		},
		{
			name:     "UPPER_CASE",
			input:    "UPPER_CASE",
			expected: "UpperCase",
		},
		{
			name:     "lowercase",
			input:    "lowercase",
			expected: "Lowercase",
		},
		{
			name:     "空字符串",
			input:    "",
			expected: "",
		},
		{
			name:     "单个字符",
			input:    "a",
			expected: "A",
		},
		{
			name:     "连续下划线",
			input:    "multi__underscore",
			expected: "MultiUnderscore",
		},
		{
			name:     "前导下划线",
			input:    "_private_field",
			expected: "PrivateField",
		},
		{
			name:     "尾随下划线",
			input:    "trailing_",
			expected: "Trailing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "驼峰转kebab",
			input:    "CamelCase",
			expected: "camel-case",
		},
		{
			name:     "连续大写",
			input:    "HTMLParser",
			expected: "html-parser",
		},
		{
			name:     "下划线转kebab",
			input:    "under_score",
			expected: "under-score",
		},
		{
			name:     "混合格式",
			input:    "Camel_Snake_HTML",
			expected: "camel-snake-html",
		},
		{
			name:     "已有连字符",
			input:    "already-hyphenated",
			expected: "already-hyphenated",
		},
		{
			name:     "数字混合",
			input:    "Version2Update",
			expected: "version2-update",
		},
		{
			name:     "全大写",
			input:    "UPPERCASE",
			expected: "uppercase",
		},
		{
			name:     "全小写",
			input:    "lowercase",
			expected: "lowercase",
		},
		{
			name:     "空字符串",
			input:    "",
			expected: "",
		},
		{
			name:     "单个字符",
			input:    "A",
			expected: "a",
		},
		{
			name:     "特殊字符",
			input:    "Test@Case#123",
			expected: "test@case#123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToKebabCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToKebabCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
