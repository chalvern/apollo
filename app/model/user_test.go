package model

import (
	"fmt"
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	u := User{
		Email:    "zhjw43@163.com",
		Password: "123456",
	}

	err := u.Create()
	assert.Nil(t, err)
}

func ftCreateOneUser() *User {
	u := &User{
		Nickname: "chalvern",
		Email:    "zhjw43@163.com",
		Password: "123456",
	}
	u.Create()
	return u
}

func TestUserUpdate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	u := ftCreateOneUser()
	u.Password = "12345678"
	err := u.Update()
	assert.Nil(t, err)
}

func TestUserFindByEmail(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()
	u1 := ftCreateOneUser()

	u2, err := u1.FindByEmail(u1.Email)
	assert.Nil(t, err)
	assert.Equal(t, u1.ID, u2.ID)
}

func TestUserFindByUID(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()
	u1 := ftCreateOneUser()

	u2, err := u1.FindByUID(u1.ID)
	assert.Nil(t, err)
	assert.Equal(t, u1.Email, u2.Email)
}

func ftCreateSomeUser(num int) (users []User) {
	for i := 0; i < num; i++ {
		u := &User{
			Email:    fmt.Sprintf("email%d@jianzhoubian.com", i),
			Password: "123456",
		}
		u.Create()
		users = append(users, *u)
	}
	return users
}
func TestUserQueryBatch(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	num := 10
	users := ftCreateSomeUser(num)
	assert.Equal(t, num, len(users))

	us, total, err := users[0].QueryBatch(0, 11)
	assert.Nil(t, err)
	assert.Equal(t, num, total)
	assert.Equal(t, num, len(us))
}
