package service

import (
	"testing"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestShareCreate(t *testing.T) {
	mydb := initializer.DB.Begin()
	model.SetMyDB(mydb)
	defer mydb.Rollback()

	share := &model.Share{
		Title:  "I am title",
		Tag:    "博客",
		UserID: 1,
	}
	assert.Equal(t, uint(0), share.ID)
	err := ShareCreate(share)
	assert.Nil(t, err)
	assert.NotEqual(t, uint(0), share.ID)
}

func TestShareUpdates(t *testing.T) {
	mydb := initializer.DB.Begin()
	model.SetMyDB(mydb)
	defer mydb.Rollback()

	share := model.FtCreateOneShare()
	share.Title = "测试修改"

	user := &model.User{}
	user1, _ := user.FindByUID(share.UserID)
	err := ShareUpdates(share, &user1)
	assert.Nil(t, err)
}
