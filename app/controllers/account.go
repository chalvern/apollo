package controllers

import (
	"net/http"

	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/apollo/tools/jwt"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// SigninGet 获取登录页面
func SigninGet(c *gin.Context) {
	c.Set(PageTitle, "登陆")
	htmlOfOk(c, "account/signin.tpl", gin.H{})
}

// SignInPost 登陆
func SignInPost(c *gin.Context) {
	c.Set(PageTitle, "登陆")
	form := struct {
		Email     string `form:"email" binding:"required,email,lenlte=50"`
		Password  string `form:"password" binding:"required,lengte=8"`
		CaptchaID string `form:"captcha_id" binding:"required"`
		Captcha   string `form:"captcha" binding:"required"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("SigninPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "account/signin.tpl", gin.H{
			FlashError: "请检查邮箱、密码、验证码内容及格式是否填写正确",
		})
		return
	}

	// 验证码校验
	if !initializer.Captcha.Verify(form.CaptchaID, form.Captcha) {
		html(c, http.StatusBadRequest, "account/signin.tpl", gin.H{
			FlashError: "验证码错误",
		})
		return
	}

	u, err := service.UserSigninByEmail(form.Email, form.Password)
	if err != nil {
		sugar.Warnf("邮箱 %s 登录失败，密码错误。 err: %v", form.Email, err)
		html(c, http.StatusBadRequest, "account/signin.tpl", gin.H{
			FlashError: "邮箱未注册或密码错误",
		})
		return
	}

	// 设置 cookie
	token, err := jwt.NewToken(map[string]interface{}{
		"email": u.Email,
	})
	if err != nil {
		sugar.Errorf("SigninPost-NewToken-err: %s", err.Error())
		return
	}
	setJustCookie(c, token)

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "登陆成功 😆😆😆",
		"Timeout":      3,
		"RedirectURL":  "/",
		"RedirectName": "主页",
	})
}

// SignupGet 获取注册页面
func SignupGet(c *gin.Context) {
	html(c, http.StatusOK, "account/signup.tpl", gin.H{
		"PageTitle": "注册",
	})
}

// SignUpPost 注册
func SignUpPost(c *gin.Context) {
	pageTitle := "注册"
	form := struct {
		Email     string `form:"email" binding:"required,email,lenlte=50"`
		Password  string `form:"password" binding:"required,lengte=8"`
		Password2 string `form:"password2" binding:"required,gtefield=Password,ltefield=Password"`
		CaptchaID string `form:"captcha_id" binding:"required"`
		Captcha   string `form:"captcha" binding:"required"`
	}{}
	// https://github.com/go-playground/validator/tree/v8.18.2
	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("SigninPost Bind form Error: %s", errs.Error())
		// errors := errs.(validator.ValidationErrors)
		html(c, http.StatusOK, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "请检查邮箱、密码、验证码内容及格式是否填写正确",
		})
		return
	}

	// 验证码校验
	if !initializer.Captcha.Verify(form.CaptchaID, form.Captcha) {
		html(c, http.StatusBadRequest, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "验证码错误",
		})
		return
	}

	if err := service.UserSignup(form.Email, form.Password); err != nil {
		html(c, http.StatusBadRequest, "account/signup.tpl", gin.H{
			"PageTitle": pageTitle,
			FlashError:  "创建用户失败，邮箱已注册",
		})
		return
	}

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "注册成功 😆😆😆",
		"Timeout":      3,
		"RedirectURL":  "/signin",
		"RedirectName": "登陆页",
	})

}

// SignOut 注销登陆
func SignOut(c *gin.Context) {
	expireCookie(c)
	html(c, http.StatusOK, "notify/success.tpl", gin.H{
		"Info":         "已注销",
		"Timeout":      3,
		"RedirectURL":  "/",
		"RedirectName": "首页",
	})
}
