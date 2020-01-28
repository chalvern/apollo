package helper

import (
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/simplate"
)

// AddFuncMap add funcMap for Simplate
func AddFuncMap() {

	// 全局
	simplate.AddFuncMap("dataFormat", DataFormatHelper)
	simplate.AddFuncMap("now_year", NowYear)

	simplate.AddFuncMap("account_normal", AccountNormalHelper)
	simplate.AddFuncMap("account_manager", AccountManagerHelper)

	simplate.AddFuncMap("account_has_share_edit_authority", AccountHasShareEditAuthority)

	simplate.AddFuncMap("pageTitle", PageTitleHelper)
	simplate.AddFuncMap("brand_title", BrandTitleHelper)
	simplate.AddFuncMap("about_content", PageSideAboutHelper)
	simplate.AddFuncMap("str_limit_length", StringLimitLengthHelper)

	simplate.AddFuncMap("create_captcha", initializer.Captcha.CreateCaptchaHTML)
	simplate.AddFuncMap("firstChar", FirstCharacterOfHelper)

}
