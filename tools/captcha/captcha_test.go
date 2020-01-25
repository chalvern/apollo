package captcha_test

import (
	"testing"
	"time"

	"github.com/chalvern/apollo/tools/captcha"
	"github.com/stretchr/testify/assert"

	"github.com/gin-contrib/cache/persistence"
)

func TestCaptcha(t *testing.T) {
	cpt := captcha.NewCaptcha("/captcha/", persistence.NewInMemoryStore(200*time.Second))
	assert.NotNil(t, cpt)
	h := cpt.CreateCaptchaHTML()
	assert.NotNil(t, h)
}
