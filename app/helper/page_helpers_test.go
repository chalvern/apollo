package helper_test

import (
	"testing"

	"github.com/chalvern/apollo/app/helper"
	"github.com/stretchr/testify/assert"
)

func TestFirstCharacterOfHelper(t *testing.T) {
	c := helper.FirstCharacterOfHelper("")
	assert.Equal(t, "U", c)

	c = helper.FirstCharacterOfHelper("jingwei")
	assert.Equal(t, "j", c)

	c = helper.FirstCharacterOfHelper("敬维")
	assert.Equal(t, "敬", c)
}
