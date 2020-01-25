package service

import (
	"github.com/chalvern/apollo/app/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	userModel = &model.User{}
)

// UserSignin 用户登陆
func UserSignin(email, password string) (model.User, error) {
	user, err := userModel.FindByEmail(email)
	return user, err
}

// UserSignup 用户注册
func UserSignup(email, password string) error {
	newUser := &model.User{
		Email: email,
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)

	return newUser.Create()
}
