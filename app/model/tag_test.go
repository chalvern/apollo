package model

import (
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestTagCreate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	tag := &Tag{
		Name:  "博客",
		Desc:  "个人博客",
		Count: 0,
	}
	err := tag.Create()
	assert.Nil(t, err)

	t.Run("Tag's name cannot be null", func(t *testing.T) {
		tag := &Tag{
			Desc:  "个人博客",
			Count: 0,
		}
		err := tag.Create()
		assert.NotNil(t, err)
	})
}

func TestTagUpdate(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	tag := FtCreateOneTag()
	tag.Name = "Blog"
	err := tag.Update()
	assert.Nil(t, err)
}

func TestTagQueryByName(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	tag := FtCreateOneTag()
	tag2, err := tag.QueryByName("博客")
	assert.Nil(t, err)
	assert.NotEqual(t, 0, tag2.ID)

	_, err = tag.QueryByName("博客Blog")
	assert.NotNil(t, err)
}

func TestTag(t *testing.T) {
	mydb = initializer.DB.Begin()
	defer mydb.Rollback()

	num := 10
	FtCreateSomeTags(num)

	tag := FtCreateOneTag()
	tags1, total, err := tag.QueryBatch(0, 10)
	assert.Nil(t, err)
	assert.Equal(t, num+1, total)
	assert.Equal(t, 10, len(tags1))
}
