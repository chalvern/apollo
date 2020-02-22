package model

import (
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestShareCreate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := Share{
		URL:    "https://jingwei.link",
		Title:  "敬维",
		Review: "A blog",
		Tag:    "博客",
	}

	err := s.Create()
	assert.Nil(t, err)
}

func ftCreateOneShare() *Share {
	user := ftCreateOneUser()
	s := &Share{
		URL:    "https://jingwei.link",
		Title:  "敬维",
		Review: "A blog",
		Tag:    "博客",
		UserID: user.ID,
	}
	s.Create()
	return s
}
func TestShareUpdate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := ftCreateOneShare()
	s.Review = "A self-blog"
	err := s.Update()
	assert.Nil(t, err)
}

func TestShareClick(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := ftCreateOneShare()
	assert.Equal(t, 0, s.ClickCount)
	err := s.Click(s.ID)
	assert.Nil(t, err)

	// 顺便把 QueryByID 以及 preload 的 User 也测一下
	s2, err := s.QueryByID(s.ID)
	assert.Nil(t, err)
	assert.Equal(t, 1, s2.ClickCount)
	assert.NotEqual(t, 0, s2.User.ID)

}

func TestShareStar(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := ftCreateOneShare()
	err := s.Star(s.ID)
	assert.Nil(t, err)

	s2, err := s.QueryByID(s.ID)
	assert.Nil(t, err)
	assert.Equal(t, 1, s2.StarCount)
}

func TestShareComment(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := ftCreateOneShare()
	err := s.Comment(s.ID)
	assert.Nil(t, err)

	s2, err := s.QueryByID(s.ID)
	assert.Nil(t, err)
	assert.Equal(t, 1, s2.CommentCount)
}

func TestShareAggregateTagCount(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	s := &Share{
		URL:    "https://jingwei.link",
		Title:  "敬维",
		Review: "A blog",
		Tag:    "博客",
	}
	s.Create()

	count, err := s.AggregateTagCount("博客")
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func TestShareQueryBatch(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	num := 10
	for i := 0; i < num; i++ {
		ftCreateOneShare()
	}

	s := ftCreateOneShare()

	shares, total, err := s.QueryBatch(0, 10, true)
	assert.Nil(t, err)
	assert.Equal(t, num+1, total)
	assert.Equal(t, 10, len(shares))
}
