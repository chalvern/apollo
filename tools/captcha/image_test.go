package captcha

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageWriteTo(t *testing.T) {
	defaultChars := []byte{0, 1, 2, 3}
	image := NewImage(defaultChars, stdWidth, stdHeight)
	assert.NotNil(t, image)
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	fmt.Println(f.Name())
	image.WriteTo(f)
}
