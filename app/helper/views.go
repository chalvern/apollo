package helper

import (
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/simplate"
)

// AddFuncMap add funcMap for Simplate
func AddFuncMap() {

	// 全局
	simplate.AddFuncMap("now_year", NowYear)
	simplate.AddFuncMap("year_month_str", MonthYearFormatHelper)
	simplate.AddFuncMap("year_date_str", DateYearFormatHelper)
	simplate.AddFuncMap("time_internal_desc", TimeInternalDesc)

	simplate.AddFuncMap("account_normal_authority", AccountNormalHelper)
	simplate.AddFuncMap("account_manager_authority", AccountManagerHelper)
	simplate.AddFuncMap("account_super_authority", AccountSuperHelper)

	simplate.AddFuncMap("account_has_share_edit_authority", AccountHasShareEditAuthority)

	simplate.AddFuncMap("page_title", PageTitleHelper)
	simplate.AddFuncMap("brand_title", BrandTitleHelper)
	simplate.AddFuncMap("about_content", PageSideAboutHelper)

	simplate.AddFuncMap("URLPathof", URLPathOfHelper)
	simplate.AddFuncMap("link", URLPathOfHelper)

	simplate.AddFuncMap("str_limit_length", StringLimitLengthHelper)
	simplate.AddFuncMap("markdown", MarkdownHelper)
	simplate.AddFuncMap("str2html", Str2html)

	simplate.AddFuncMap("create_captcha", initializer.Captcha.CreateCaptchaHTML)
	simplate.AddFuncMap("firstChar", FirstCharacterOfHelper)

}
