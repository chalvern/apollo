package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountValidEmailBody(t *testing.T) {
	body, err := accountValidEmailBody("jianzhoubian@163.com", "Jingwei", "iamtoken")
	assert.Nil(t, err)
	assert.Contains(t, body, "验证邮箱")
}

// 此测试用例仅用来手动测试，因为它真的会发邮件出去
func testAccountValidEmail(t *testing.T) {
	err := AccountValidEmail("jianzhoubian@163.com", "Jingwei", "iamtoken")
	assert.Nil(t, err)
}
