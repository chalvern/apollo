package helper

import "time"

// NowYear 当前的年份
func NowYear() string {
	return time.Now().Format("2006")
}

// MonthYearFormatHelper 返回月份+年份的日期格式
// 提供类似 2006年01月02日03时04分05秒UTC 的模板
func MonthYearFormatHelper(t time.Time) string {
	return t.Format("2006-01")
}

// DateYearFormatHelper 返回日期+年份的日期格式
func DateYearFormatHelper(t time.Time) string {
	return t.Format("2006-01-02")
}

// TimeInternalDesc 时间间隔描述
func TimeInternalDesc(t time.Time) string {
	now := time.Now()
	internalHour := now.Sub(t).Hours()
	switch {
	case internalHour <= 24:
		return "刚刚"
	case internalHour <= 24*30:
		return "1天前"
	case internalHour <= 24*30*3:
		return "1个月前"
	case internalHour <= 24*30*6:
		return "3个月前"
	default:
		return "1年前"
	}
}
