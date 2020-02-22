package model

import (
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestCommentCreate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	comment := Comment{}
	err := comment.Create()
	assert.Nil(t, err)
}

func TestCommentUpdate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	comment := FtCreateOneComment()
	comment.Reply = "so boring"
	err := comment.Update()
	assert.Nil(t, err)
}
