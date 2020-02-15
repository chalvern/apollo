package service

import (
	"fmt"
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
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
func UserSigninByEmail(email, password string) (*model.User, error) {
	user, err := userModel.FindByEmail(email)
	// password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &user, err
	}
	return &user, nil
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
func UserSignup(email, password, nickname string) (*model.User, error) {
	newUser := &model.User{
		Email: email,
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser.Password = string(hash)
	newUser.Nickname = nickname
	newUser.EmailValidToken = uuid.Must(uuid.NewV4()).String()

	// 查看是否为超级用户
	if viper.GetString("admin.super") == email {
		newUser.Priority = ^0
	}

	err = newUser.Create()
	return newUser, err
}

// UserUpdates 更新用户
func UserUpdates(user *model.User) error {
	if user.ID == 0 {
		return fmt.Errorf("更新分享必须是已存在的用户")
	}
	return userUpdates(user)
}

// userUpdates 更新用户
func userUpdates(user *model.User) error {
	return user.Update()
}

// UserValidEmail 用户验证邮箱
func UserValidEmail(email, token string) error {
	user, err := userModel.FindByEmail(email)
	if err != nil {
		return err
	}
	if user.EmailValidToken != token {
		return fmt.Errorf("传入的 token 无效")
	}
	if user.EmailVarified {
		return fmt.Errorf("用户邮箱 %s 已认证", email)
	}
	// 到这里说明用户身份正确
	user.EmailVarified = true
	user.Priority = user.Priority | model.UserPriorityCommon
	return UserUpdates(&user)
}
