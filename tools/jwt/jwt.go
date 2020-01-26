package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	hmacSecret     = []byte("123456")
	expDuration, _ = time.ParseDuration("2400h")
	isInitialed    = false
)

// SetHmacSecret 设置 hmac 的密码
func SetHmacSecret(secret string) {
	hmacSecret = []byte(secret)
	isInitialed = true
}

// SetExpDuration 配置超时时间
func SetExpDuration(exp time.Duration) {
	expDuration = exp
}

// NewToken 新创建一个 token
func NewToken(mcTemp map[string]interface{}) (string, error) {
	if !isInitialed {
		fmt.Println("jwt's Secret must be set before using.")
	}
	// copy
	mc := jwt.MapClaims{}
	for k, v := range mcTemp {
		mc[k] = v
	}
	// 配置过期时间
	if mc["exp"] == nil {
		mc["exp"] = time.Now().Add(expDuration).Unix()
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(hmacSecret)
}

// ParseToken 解析
func ParseToken(tokenString string) (jwt.MapClaims, error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}
