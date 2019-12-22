package helper

import (
	"github.com/chalvern/simplate"
)

// AddFuncMap add funcMap for Simplate
func AddFuncMap() {

	// 全局
	simplate.AddFuncMap("dataFormat", DataFormatHelper)
	simplate.AddFuncMap("pageTitle", PageTitleHelper)
	simplate.AddFuncMap("firstChar", FirstCharacterOfHelper)
}
