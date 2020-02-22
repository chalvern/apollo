package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chalvern/apollo/app/controllers"
	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/apollo/tools/jwt"
	"github.com/stretchr/testify/assert"
)

func TestHomeIndex(t *testing.T) {
	mydb := initializer.DB.Begin()
	model.SetMyDB(mydb)
	defer mydb.Rollback()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	rR.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	bodyContent := w.Body.String()
	// fmt.Print(bodyContent)
	assert.Contains(t, bodyContent, ">按标签</a>")
	assert.Contains(t, bodyContent, "登录")
	assert.Contains(t, bodyContent, "注册")

	t.Run("home index with login", func(t *testing.T) {
		// 设置 cookie
		user := model.FtCreateOneUser()
		token, _ := jwt.NewToken(map[string]interface{}{
			"email": user.Email,
		})

		req, _ := http.NewRequest("GET", "/", nil)
		req.AddCookie(&http.Cookie{
			Name:  controllers.CookieTag,
			Value: token,
		})
		w := httptest.NewRecorder()

		rR.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		bodyContent1 := w.Body.String()
		assert.NotContains(t, bodyContent1, "登录")
		assert.Contains(t, bodyContent1, "个人资料")
	})
}

func TestHomeAboutHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about", nil)
	rR.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	bodyContent := w.Body.String()
	assert.Contains(t, bodyContent, "关于")
}
