package goutil

import (
	"errors"
	"time"
)

// DateNumber 获取当前日期数字（格式：20060102）
func DateNumber() int64 {
	now := time.Now()
	return int64(now.Year()*10000 + int(now.Month())*100 + now.Day())
}

// DayStartTime 获取今日开始时间（00:00:00）
func DayStartTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// DayEndTime 获取今日结束时间（23:59:59）
func DayEndTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

// YesterdayStartTime 获取昨日开始时间
func YesterdayStartTime() time.Time {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, now.Location())
}

// YesterdayEndTime 获取昨日结束时间
func YesterdayEndTime() time.Time {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0, now.Location())
}

// TomorrowStartTime 获取明日开始时间
func TomorrowStartTime() time.Time {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	return time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, now.Location())
}

// IsSameDay 判断两个时间是否为同一天
func IsSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day()
}

// DaysBetween 计算两个日期之间相隔的天数
func DaysBetween(start, end time.Time) int {
	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	endDate := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	return int(endDate.Sub(startDate).Hours() / 24)
}

// BeginningOfWeek 获取本周开始时间（周一 00:00:00）
func BeginningOfWeek(t time.Time) time.Time {
	// 计算到周一需要减去的天数
	offset := int(t.Weekday())
	if offset == 0 { // 周日
		offset = 6
	} else {
		offset -= 1
	}

	weekStart := t.AddDate(0, 0, -offset)
	return time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfWeek 获取本周结束时间（周日 23:59:59）
func EndOfWeek(t time.Time) time.Time {
	// 计算到周日需要增加的天数
	offset := int(t.Weekday())
	if offset == 0 { // 周日
		offset = 0
	} else {
		offset = 7 - offset
	}

	weekEnd := t.AddDate(0, 0, offset)
	return time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 0, t.Location())
}

// GetMonthStartEnd 获取指定时间所在月份的开始和结束时间
func GetMonthStartEnd(t time.Time) (time.Time, time.Time, error) {
	if t.IsZero() {
		return time.Time{}, time.Time{}, errors.New("输入时间不能为空")
	}

	// 获取月份第一天 00:00:00
	monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// 获取下个月第一天，然后减一秒得到本月最后一天 23:59:59
	nextMonth := monthStart.AddDate(0, 1, 0)
	monthEnd := nextMonth.Add(-time.Second)

	return monthStart, monthEnd, nil
}
