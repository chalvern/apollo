package helper

import "github.com/spf13/viper"
import "github.com/chalvern/apollo/configs/constants"

// PageTitleHelper 组装得到和页面 title
func PageTitleHelper(title string) string {
	pageTitle := viper.GetString(constants.PageTitle)
	return title + "-" + pageTitle
}

// FirstCharacterOfHelper 获取字符串的第一个字符
func FirstCharacterOfHelper(name string) string {
	if len(name) == 0 {
		return "U"
	}
	return name[:1]
}
