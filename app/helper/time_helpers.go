package helper

import "time"

// DataFormatHelper 格式化时间为固定格式字符串
func DataFormatHelper(t time.Time) string {
	return t.UTC().Format("2006年01月02日03时04分05秒UTC")
}

// NowYear 当前的年份
func NowYear() string {
	return time.Now().Format("2006")
}
