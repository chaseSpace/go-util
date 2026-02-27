package ustr

import "testing"

func TestCamelToUnderScore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"基本驼峰", "CamelCase", "camel_case"},
		{"简单单词", "User", "user"},
		{"多个大写字母", "HTMLParser", "html_parser"},
		{"连续大写", "XMLHttpRequest", "xml_http_request"},
		{"数字混合", "UserID2Name", "user_id2_name"},
		{"_already_snake", "already_snake", "already_snake"},
		{"空字符串", "", ""},
		{"单个字符", "A", "a"},
		{"全大写", "UPPERCASE", "uppercase"},
		{"全小写", "lowercase", "lowercase"},
		{"复杂案例", "iPhoneXSMax", "i_phone_xs_max"},
		{"带数字开头", "Version2Update", "version2_update"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CamelToUnderScore(tt.input)
			if result != tt.expected {
				t.Errorf("CamelToUnderScore(%q) = %q, 期望 %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUnderScoreToCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"基本下划线", "camel_case", "CamelCase"},
		{"简单单词", "user", "User"},
		{"多个下划线", "html_parser", "HtmlParser"},
		{"连续下划线", "xml_http_request", "XmlHttpRequest"},
		{"带数字", "user_id2_name", "UserId2Name"},
		{"AlreadyCamel", "AlreadyCamel", "Alreadycamel"},
		{"空字符串", "", ""},
		{"单个字符", "a", "A"},
		{"全大写下划线", "UPPER_CASE", "UpperCase"},
		{"全小写下划线", "lower_case", "LowerCase"},
		{"复杂案例", "i_phone_xs_max", "IPhoneXsMax"},
		{"前导下划线", "_private_field", "PrivateField"},
		{"尾随下划线", "trailing_", "Trailing"},
		{"多重下划线", "multi___underscore", "MultiUnderscore"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnderScoreToCamel(tt.input)
			if result != tt.expected {
				t.Errorf("UnderScoreToCamel(%q) = %q, 期望 %q", tt.input, result, tt.expected)
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
