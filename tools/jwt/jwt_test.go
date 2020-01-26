package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	mc := map[string]interface{}{
		"uid": uint(1),
	}
	token, err := NewToken(mc)
	assert.Nil(t, err)
	// parse
	m, err := ParseToken(token)
	i := m["uid"].(float64)
	assert.Nil(t, err)
	assert.Equal(t, mc["uid"], uint(i))
	assert.Nil(t, m.Valid())
}

func TestNewTokenValid(t *testing.T) {

	var expDuration, _ = time.ParseDuration("-24h")

	mc := map[string]interface{}{
		"uid": uint(1),
		"exp": time.Now().Add(expDuration).Unix(),
	}
	token, err := NewToken(mc)
	assert.Nil(t, err)
	// parse
	_, err = ParseToken(token)
	assert.NotNil(t, err)
}
