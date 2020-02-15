package helper

import (
	"github.com/chalvern/apollo/configs/constants"
	"github.com/spf13/viper"
)

// PageTitleHelper 组装得到和页面 title
func PageTitleHelper(title string) string {
	pageTitle := viper.GetString(constants.PageTitle)
	return title + "-" + pageTitle
}

// PageSideAboutHelper 边框中的介绍
func PageSideAboutHelper() string {
	aboutContent := viper.GetString(constants.AboutContent)
	return aboutContent
}

// BrandTitleHelper 返回 logo 描述
func BrandTitleHelper() string {
	brandTitle := viper.GetString(constants.BrandTitle)
	if brandTitle == "" {
		return "applo"
	}
	return brandTitle
}

// FirstCharacterOfHelper 获取字符串的第一个字符
func FirstCharacterOfHelper(name string) string {
	if len(name) == 0 {
		return "U"
	}

	rawRune := []rune(name)
	return string(rawRune[:1])
}
