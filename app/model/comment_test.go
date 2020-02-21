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

// for test
func ftCreateOneComment() *Comment {
	comment := &Comment{
		UserID: 1,
		Reply:  "Boring",
	}
	comment.Create()
	return comment
}

func TestCommentUpdate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	comment := ftCreateOneComment()
	comment.Reply = "so boring"
	err := comment.Update()
	assert.Nil(t, err)
}
