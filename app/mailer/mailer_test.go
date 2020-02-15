package mailer

import (
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	initializer.InitViperWithFile("../../configs/config.yml")
	Init()

	mailTo := []string{"jianzhoubian@163.com"}
	err := sendMail(mailTo, "subject", "body")
	assert.Nil(t, err)
}
