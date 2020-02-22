package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 此测试用例仅用来手动测试，因为它真的会发邮件出去
func testSendMail(t *testing.T) {
	mailTo := []string{"jianzhoubian@163.com"}
	err := sendMail(mailTo, "subject", "body")
	assert.Nil(t, err)
}
