package service

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGoResty(t *testing.T) {
	// Create a Resty Client
	client := resty.New()

	// just mentioning about POST as an example with simple flow
	// User Login
	_, err := client.R().
		SetFormData(map[string]string{
			"username": "jeeva",
			"password": "mypass",
		}).
		Post("http://localhost:2020/login")
	assert.NotNil(t, err)

}
