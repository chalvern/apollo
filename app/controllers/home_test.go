package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeIndex(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	rR.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	bodyContent := w.Body.String()
	// fmt.Print(bodyContent)
	assert.Contains(t, bodyContent, ">按标签</a>")
}

func TestHomeAboutHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about", nil)
	rR.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	bodyContent := w.Body.String()
	assert.Contains(t, bodyContent, "关于")
}
