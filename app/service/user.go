package service

import (
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	userModel = &model.User{}
)

// UsersQueryWithContext 根据 url query 参数检索
func UsersQueryWithContext(c *gin.Context, args ...interface{}) (users []model.User, allPage int, err error) {

	page := queryPage(c)
	pageSize := queryPageSize(c)
	argS, argArray := argsInit(args...)
	argArray[0] = strings.Join(argS, "AND")

	offset := (page - 1) * pageSize
	users, total, err := userModel.QueryBatch(offset, pageSize, args...)
	allPage = total/pageSize + 1
	return
}

// UserSigninByEmail 用户登陆
func UserSigninByEmail(email, password string) (model.User, error) {
	user, err := userModel.FindByEmail(email)
	// password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

// UserFindByEmail 根据邮件查询用户
func UserFindByEmail(email string) (*model.User, error) {
	user, err := userModel.FindByEmail(email)
	return &user, err
}

// UserFindByUID 根据 id 检索用户
func UserFindByUID(uid interface{}) (*model.User, error) {
	user, err := userModel.FindByUID(uid)
	return &user, err
}

// UserSignup 用户注册
func UserSignup(email, password, nickname string) error {
	newUser := &model.User{
		Email: email,
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)
	newUser.Nickname = nickname

	return newUser.Create()
}
