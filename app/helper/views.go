package helper

import (
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/simplate"
)

// AddFuncMap add funcMap for Simplate
func AddFuncMap() {

	// 全局
	simplate.AddFuncMap("dataFormat", DataFormatHelper)
	simplate.AddFuncMap("pageTitle", PageTitleHelper)
	simplate.AddFuncMap("brand_title", BrandTitleHelper)
	simplate.AddFuncMap("about_content", PageSideAboutHelper)

	simplate.AddFuncMap("create_captcha", initializer.Captcha.CreateCaptchaHTML)
	simplate.AddFuncMap("firstChar", FirstCharacterOfHelper)

}
