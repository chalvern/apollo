package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	mailTo := []string{"jianzhoubian@163.com"}
	err := sendMail(mailTo, "subject", "body")
	assert.Nil(t, err)
}
