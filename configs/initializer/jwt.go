package initializer

import (
	"time"

	"github.com/chalvern/apollo/tools/jwt"
	"github.com/spf13/viper"
)

// InitJwt 初始化 jwt
func InitJwt() {
	jwt.SetExpDuration(time.Hour * 24 * 100)
	jwt.SetHmacSecret(viper.GetString("jwt.hmac_secret"))
}
