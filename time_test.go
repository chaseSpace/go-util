package go_util

import (
	"testing"
	"time"
)

func TestDateNumber(t *testing.T) {
	now := time.Now()
	expected := int64(now.Year()*10000 + int(now.Month())*100 + now.Day())
	result := DateNumber()

	if result != expected {
		t.Errorf("DateNumber() = %d, 期望 %d", result, expected)
	}

	// 验证格式正确性
	dateStr := time.Now().Format("20060102")
	expectedFromStr := int64(0)
	for _, digit := range dateStr {
		expectedFromStr = expectedFromStr*10 + int64(digit-'0')
	}

	if result != expectedFromStr {
		t.Errorf("DateNumber格式不正确，期望 %d，实际 %d", expectedFromStr, result)
	}
}

func TestDayStartTime(t *testing.T) {
	result := DayStartTime()
	now := time.Now()

	// 验证日期正确
	if result.Year() != now.Year() ||
		result.Month() != now.Month() ||
		result.Day() != now.Day() {
		t.Errorf("日期不正确: 期望 %d-%02d-%02d，实际 %d-%02d-%02d",
			now.Year(), now.Month(), now.Day(),
			result.Year(), result.Month(), result.Day())
	}

	// 验证时间为00:00:00
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("时间不正确: 期望 00:00:00，实际 %02d:%02d:%02d",
			result.Hour(), result.Minute(), result.Second())
	}
}

func TestDayEndTime(t *testing.T) {
	result := DayEndTime()
	now := time.Now()

	// 验证日期正确
	if result.Year() != now.Year() ||
		result.Month() != now.Month() ||
		result.Day() != now.Day() {
		t.Errorf("日期不正确")
	}

	// 验证时间为23:59:59
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("时间不正确: 期望 23:59:59，实际 %02d:%02d:%02d",
			result.Hour(), result.Minute(), result.Second())
	}
}

func TestYesterdayAndTomorrow(t *testing.T) {
	now := time.Now()
	yesterdayStart := YesterdayStartTime()
	yesterdayEnd := YesterdayEndTime()
	tomorrowStart := TomorrowStartTime()

	// 验证昨天开始时间
	expectedYesterday := now.AddDate(0, 0, -1)
	if yesterdayStart.Year() != expectedYesterday.Year() ||
		yesterdayStart.Month() != expectedYesterday.Month() ||
		yesterdayStart.Day() != expectedYesterday.Day() {
		t.Errorf("昨天开始时间日期不正确")
	}

	// 验证昨天结束时间和今天开始时间的关系
	if yesterdayEnd.After(DayStartTime()) {
		t.Errorf("昨天结束时间不应晚于今天开始时间")
	}

	// 验证明天开始时间
	expectedTomorrow := now.AddDate(0, 0, 1)
	if tomorrowStart.Year() != expectedTomorrow.Year() ||
		tomorrowStart.Month() != expectedTomorrow.Month() ||
		tomorrowStart.Day() != expectedTomorrow.Day() {
		t.Errorf("明天开始时间日期不正确")
	}

	// 验证时间顺序
	if !yesterdayStart.Before(DayStartTime()) ||
		!DayStartTime().Before(tomorrowStart) {
		t.Errorf("时间顺序不正确")
	}
}

func TestIsSameDay(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{"同一天不同时刻", now, now.Add(1 * time.Hour), true},
		{"不同天", now, now.AddDate(0, 0, 1), false},
		{"同一天跨月",
			time.Date(2026, 2, 28, 23, 59, 59, 0, time.Local),
			time.Date(2026, 3, 1, 0, 0, 0, 0, time.Local),
			false},
		{"完全相同", now, now, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameDay(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameDay(%v, %v) = %v, 期望 %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{"同一天", now, now, 0},
		{"相差1天", now, now.AddDate(0, 0, 1), 1},
		{"相差5天", now, now.AddDate(0, 0, 5), 5},
		{"反向", now.AddDate(0, 0, 5), now, -5},
		{"跨月",
			time.Date(2026, 1, 31, 10, 0, 0, 0, time.Local),
			time.Date(2026, 2, 2, 15, 0, 0, 0, time.Local),
			2},
		{"跨年",
			time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local),
			time.Date(2026, 1, 2, 0, 0, 0, 0, time.Local),
			2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DaysBetween(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("DaysBetween() = %d, 期望 %d", result, tt.expected)
			}
		})
	}
}

func TestBeginningAndEndOfWeek(t *testing.T) {
	// 测试一周的每一天
	for i := 0; i < 7; i++ {
		testTime := time.Now().AddDate(0, 0, i)
		weekStart := BeginningOfWeek(testTime)
		weekEnd := EndOfWeek(testTime)

		// 验证周开始是周一
		if weekStart.Weekday() != time.Monday {
			t.Errorf("周开始应该是周一，实际是 %s", weekStart.Weekday())
		}

		// 验证周结束是周日
		if weekEnd.Weekday() != time.Sunday {
			t.Errorf("周结束应该是周日，实际是 %s", weekEnd.Weekday())
		}

		// 验证时间正确
		if weekStart.Hour() != 0 || weekStart.Minute() != 0 || weekStart.Second() != 0 {
			t.Errorf("周开始时间应该是 00:00:00")
		}

		if weekEnd.Hour() != 23 || weekEnd.Minute() != 59 || weekEnd.Second() != 59 {
			t.Errorf("周结束时间应该是 23:59:59")
		}

		// 验证周期正确（7天）
		duration := weekEnd.Sub(weekStart)
		expectedDuration := 7*24*time.Hour - 1*time.Second
		if duration != expectedDuration {
			t.Errorf("周周期应该是7天，实际 %v", duration)
		}
	}
}
