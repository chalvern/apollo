package model

import "fmt"

// FtCreateOneUser ft
func FtCreateOneUser() *User {
	u := &User{
		Nickname: "chalvern",
		Email:    "zhjw43@163.com",
		Password: "123456",
	}
	u.Create()
	return u
}

// FtCreateSomeUser ft
func FtCreateSomeUser(num int) (users []User) {
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

// FtCreateOneComment for test
func FtCreateOneComment() *Comment {
	comment := &Comment{
		UserID: 1,
		Reply:  "Boring",
	}
	comment.Create()
	return comment
}

// FtCreateOneShare for test
func FtCreateOneShare() *Share {
	user := FtCreateOneUser()
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

// FtCreateOneTag ft
func FtCreateOneTag() *Tag {
	tag := &Tag{
		Name:  "博客",
		Desc:  "个人博客",
		Count: 0,
	}
	tag.Create()
	return tag
}

// FtCreateSomeTags fot test
func FtCreateSomeTags(num int) (tags []Tag) {
	for i := 0; i < num; i++ {
		tag := &Tag{
			Name:  fmt.Sprintf("博客%d", i),
			Desc:  "个人博客",
			Count: 0,
		}
		tag.Create()
		tags = append(tags, *tag)
	}
	return tags
}
