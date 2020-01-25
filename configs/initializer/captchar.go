package initializer

import (
	"context"
	"time"

	"github.com/gin-contrib/cache/persistence"

	"github.com/chalvern/apollo/tools/captcha"
)

var (
	// Captcha 验证码
	Captcha *captcha.Captcha
)

// InitCaptcha 初始化验证码
func InitCaptcha(ctx context.Context) {
	Captcha = captcha.NewCaptcha("/captcha/", persistence.NewInMemoryStore(200*time.Second))
}
