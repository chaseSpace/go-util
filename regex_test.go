package go_util

import (
	"testing"
	"time"
)

func TestIsIP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效IPv4", "192.168.1.1", true},
		{"有效IPv4边界", "255.255.255.255", true},
		{"有效IPv4最小值", "0.0.0.0", true},
		{"无效IPv4超出范围", "256.1.1.1", false},
		{"无效IPv4负数", "192.168.-1.1", false},
		{"无效IPv4缺少段", "192.168.1", false},
		{"无效IPv4多段", "192.168.1.1.1", false},
		{"空字符串", "", false},
		{"IPv6地址", "2001:db8::1", false},
		{"域名", "www.example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIP(tt.input)
			if result != tt.expected {
				t.Errorf("IsIP(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效IPv4", "192.168.1.1", true},
		{"本地回环", "127.0.0.1", true},
		{"IPv6", "2001:db8::1", false},
		{"域名", "localhost", false},
		{"空字符串", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIPv4(tt.input)
			if result != tt.expected {
				t.Errorf("IsIPv4(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"完整IPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"压缩IPv6", "2001:db8::1", true},
		{"本地回环IPv6", "::1", true},
		{"全零IPv6", "::", true},
		{"IPv4", "192.168.1.1", false},
		{"无效IPv6", "2001:db8:::1", false},
		{"空字符串", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIPv6(tt.input)
			if result != tt.expected {
				t.Errorf("IsIPv6(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsChinaPhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效手机号", "13812345678", true},
		{"移动号码", "13912345678", true},
		{"联通号码", "13012345678", true},
		{"电信号码", "13312345678", true},
		{"新号段", "16612345678", true},
		{"无效前缀", "12812345678", false},
		{"长度不足", "1381234567", false},
		{"长度过长", "138123456789", false},
		{"包含字母", "138abc45678", false},
		{"空字符串", "", false},
		{"固定电话", "010-12345678", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsChinaPhone(tt.input)
			if result != tt.expected {
				t.Errorf("IsChinaPhone(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"标准邮箱", "test@example.com", true},
		{"带点的用户名", "test.user@example.com", true},
		{"带加号", "test+tag@example.com", true},
		{"带下划线", "test_user@example.com", true},
		{"子域名", "user@mail.example.com", true},
		{"多级域名", "user@example.co.uk", true},
		{"缺少@符号", "testexample.com", false},
		{"缺少域名", "test@", false},
		{"缺少用户名", "@example.com", false},
		{"无效字符", "test@exam ple.com", false},
		{"空字符串", "", false},
		{"只有@", "@", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEmail(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmail(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"HTTP网址", "http://example.com", true},
		{"HTTPS网址", "https://example.com", true},
		{"FTP网址", "ftp://example.com", true},
		{"带路径", "https://example.com/path", true},
		{"带查询参数", "https://example.com?param=value", true},
		{"缺少协议", "example.com", false},
		{"无效协议", "htp://example.com", false},
		{"空字符串", "", false},
		{"只有协议", "http://", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsURL(tt.input)
			if result != tt.expected {
				t.Errorf("IsURL(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsChinaIDCard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效身份证", "110101199001011234", true},
		{"有效身份证2", "440524188001010014", true},
		{"长度不足", "11010119900101123", false},
		{"长度过长", "1101011990010112345", false},
		{"包含字母", "11010119900101123a", false},
		{"无效地区码", "000101199001011234", false},
		{"无效日期", "110101199013011234", false},
		{"空字符串", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsChinaIDCard(tt.input)
			if result != tt.expected {
				t.Errorf("IsChinaIDCard(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaskPhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"有效手机号", "13812345678", "138****5678"},
		{"无效手机号", "12345678901", "12345678901"},
		{"空字符串", "", ""},
		{"长度不足", "1381234567", "1381234567"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskPhone(tt.input)
			if result != tt.expected {
				t.Errorf("MaskPhone(%q) = %q, 期望 %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"标准邮箱", "test@example.com", "t***t@example.com"},
		{"短用户名", "ab@example.com", "ab@example.com"},
		{"无效邮箱", "invalid-email", "invalid-email"},
		{"空字符串", "", ""},
		{"复杂用户名", "test.user+tag@example.com", "t***g@example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskEmail(tt.input)
			if result != tt.expected {
				t.Errorf("MaskEmail(%q) = %q, 期望 %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetMonthStartEnd(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		wantStart time.Time
		wantEnd   time.Time
		wantErr   bool
	}{
		{
			name:      "正常月份",
			input:     time.Date(2026, 2, 15, 10, 30, 45, 0, time.Local),
			wantStart: time.Date(2026, 2, 1, 0, 0, 0, 0, time.Local),
			wantEnd:   time.Date(2026, 2, 28, 23, 59, 59, 0, time.Local),
			wantErr:   false,
		},
		{
			name:      "一月",
			input:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local),
			wantStart: time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local),
			wantEnd:   time.Date(2026, 1, 31, 23, 59, 59, 0, time.Local),
			wantErr:   false,
		},
		{
			name:      "十二月",
			input:     time.Date(2026, 12, 31, 23, 59, 59, 0, time.Local),
			wantStart: time.Date(2026, 12, 1, 0, 0, 0, 0, time.Local),
			wantEnd:   time.Date(2026, 12, 31, 23, 59, 59, 0, time.Local),
			wantErr:   false,
		},
		{
			name:      "闰年二月",
			input:     time.Date(2024, 2, 15, 12, 0, 0, 0, time.Local),
			wantStart: time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local),
			wantEnd:   time.Date(2024, 2, 29, 23, 59, 59, 0, time.Local),
			wantErr:   false,
		},
		{
			name:      "空时间",
			input:     time.Time{},
			wantStart: time.Time{},
			wantEnd:   time.Time{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, err := GetMonthStartEnd(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetMonthStartEnd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !gotStart.Equal(tt.wantStart) {
				t.Errorf("GetMonthStartEnd() 开始时间 = %v, 期望 %v", gotStart, tt.wantStart)
			}

			if !gotEnd.Equal(tt.wantEnd) {
				t.Errorf("GetMonthStartEnd() 结束时间 = %v, 期望 %v", gotEnd, tt.wantEnd)
			}

			// 验证开始时间是当月1号 00:00:00
			if !tt.wantErr && (gotStart.Day() != 1 ||
				gotStart.Hour() != 0 ||
				gotStart.Minute() != 0 ||
				gotStart.Second() != 0) {
				t.Errorf("开始时间格式不正确: %v", gotStart)
			}

			// 验证结束时间是月末 23:59:59
			if !tt.wantErr && (gotEnd.Hour() != 23 ||
				gotEnd.Minute() != 59 ||
				gotEnd.Second() != 59) {
				t.Errorf("结束时间格式不正确: %v", gotEnd)
			}
		})
	}
}
