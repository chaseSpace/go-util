package go_util

import (
	"testing"
	"unicode"
)

func TestRandStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"正常长度", 10},
		{"零长度", 0},
		{"负长度", -5},
		{"大长度", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStr(tt.length)

			// 验证长度
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("期望空字符串，但得到: %s", result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("期望长度 %d，但得到 %d", tt.length, len(result))
			}

			// 验证字符类型（应只包含字母和数字）
			for _, char := range result {
				if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
					t.Errorf("字符串包含非法字符: %c", char)
				}
			}
		})
	}
}

func TestRandStrDigit(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"正常长度", 8},
		{"零长度", 0},
		{"负长度", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStrDigit(tt.length)

			// 验证长度
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("期望空字符串，但得到: %s", result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("期望长度 %d，但得到 %d", tt.length, len(result))
			}

			// 验证只包含数字
			for _, char := range result {
				if !unicode.IsDigit(char) {
					t.Errorf("字符串包含非数字字符: %c", char)
				}
			}
		})
	}
}

func TestRandStrAlpha(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"正常长度", 12},
		{"零长度", 0},
		{"负长度", -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStrLetter(tt.length)

			// 验证长度
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("期望空字符串，但得到: %s", result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("期望长度 %d，但得到 %d", tt.length, len(result))
			}

			// 验证只包含字母
			for _, char := range result {
				if !unicode.IsLetter(char) {
					t.Errorf("字符串包含非字母字符: %c", char)
				}
			}
		})
	}
}

func TestRandStrEmoji(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"正常长度", 5},
		{"零长度", 0},
		{"负长度", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStrEmoji(tt.length)

			// 验证长度
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("期望空字符串，但得到: %s", result)
				}
				return
			}

			runes := []rune(result)
			if len(runes) != tt.length {
				t.Errorf("期望长度 %d，但得到 %d", tt.length, len(runes))
			}
		})
	}
}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"正常范围", 1, 10},
		{"相同值", 5, 5},
		{"逆序范围", 10, 1},
		{"负数范围", -10, -1},
		{"跨零范围", -5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandInt(tt.min, tt.max)

			// 验证边界情况
			if tt.min >= tt.max {
				if result != tt.min {
					t.Errorf("当min>=max时，期望返回min(%d)，但得到%d", tt.min, result)
				}
				return
			}

			// 验证范围
			if result < tt.min || result >= tt.max {
				t.Errorf("结果%d不在范围[%d, %d)内", result, tt.min, tt.max)
			}
		})
	}
}

func TestRandFloat(t *testing.T) {
	tests := []struct {
		name string
		min  float64
		max  float64
	}{
		{"正常范围", 0.0, 1.0},
		{"相同值", 3.14, 3.14},
		{"逆序范围", 5.0, 1.0},
		{"负数范围", -10.0, -5.0},
		{"跨零范围", -2.5, 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandFloat(tt.min, tt.max)

			// 验证边界情况
			if tt.min >= tt.max {
				if result != tt.min {
					t.Errorf("当min>=max时，期望返回min(%f)，但得到%f", tt.min, result)
				}
				return
			}

			// 验证范围
			if result < tt.min || result >= tt.max {
				t.Errorf("结果%f不在范围[%f, %f)内", result, tt.min, tt.max)
			}
		})
	}
}

func TestRandChoice(t *testing.T) {
	// 测试字符串切片
	strSlice := []string{"apple", "banana", "cherry", "date"}
	result := RandChoice(strSlice)

	// 验证结果在切片中
	found := false
	for _, item := range strSlice {
		if item == result {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("选择的结果%s不在原切片中", result)
	}

	// 测试空切片
	emptySlice := []string{}
	emptyResult := RandChoice(emptySlice)
	if emptyResult != "" {
		t.Errorf("空切片应该返回零值，但得到: %s", emptyResult)
	}

	// 测试整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	intResult := RandChoice(intSlice)
	valid := false
	for _, item := range intSlice {
		if item == intResult {
			valid = true
			break
		}
	}
	if !valid {
		t.Errorf("选择的整数%d不在原切片中", intResult)
	}
}

func TestRandomness(t *testing.T) {
	// 测试生成的随机字符串确实不同
	results := make(map[string]bool)
	for i := 0; i < 100; i++ {
		result := RandStr(10)
		if results[result] {
			t.Errorf("生成了重复的随机字符串: %s", result)
		}
		results[result] = true
	}

	// 测试随机数分布
	counts := make(map[int]int)
	for i := 0; i < 1000; i++ {
		num := RandInt(1, 6) // 模拟骰子
		counts[num]++
	}

	// 验证每个数字都有出现
	for i := 1; i <= 5; i++ {
		if counts[i] == 0 {
			t.Errorf("数字%d在1000次试验中未出现", i)
		}
	}
}
