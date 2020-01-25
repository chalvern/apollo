package rand_test

import (
	"testing"

	"github.com/chalvern/apollo/tools/rand"
	"github.com/stretchr/testify/assert"
)

func TestRandomCreateBytes(t *testing.T) {
	bytes := rand.RandomCreateBytes(14)
	assert.NotNil(t, bytes)
	assert.Equal(t, 14, len(bytes))
}
