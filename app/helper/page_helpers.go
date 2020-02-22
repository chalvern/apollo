package helper

import (
	"fmt"

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

// URLPathOfHelper 根据 URL 名称转 path 地址
// params 按照 name=value 的方式分别传入 name 和 value
// 且必须成对出现
func URLPathOfHelper(name string, params ...interface{}) string {
	urlPath := routerConfig.GetAbsoluteURLOf(name)

	// 目前只接受 2 对参数
	if paramsLen := len(params); paramsLen == 2 {
		urlPath = fmt.Sprintf("%s?%v=%v", urlPath, params[0], params[1])
	} else if paramsLen == 4 {
		urlPath = fmt.Sprintf("%s?%v=%v&%v=%v", urlPath, params[0], params[1], params[2], params[3])
	}
	return urlPath
}
