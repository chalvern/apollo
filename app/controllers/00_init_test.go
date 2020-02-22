package controllers_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/chalvern/apollo/app/controllers"
	"github.com/chalvern/apollo/app/helper"
	"github.com/chalvern/apollo/app/mailer"
	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/pubsub"
	"github.com/chalvern/apollo/app/router"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/apollo/tools/validator"
	"github.com/chalvern/simplate"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	rR *gin.Engine
)

func TestMain(m *testing.M) {
	fmt.Println("begin test of controllers")

	// 加载测试环境的 yaml
	initializer.InitViperWithFile("../../configs/config_test.yml")
	initializer.InitMysql(context.Background())
	initializer.InitCaptcha(context.Background())

	// 1.1 初始化 model/controller 等
	model.Init()
	controllers.Init()
	pubsub.Init()
	mailer.Init()

	validator.InitValidatorEnhancement()

	r := gin.Default()

	_, file, _, _ := runtime.Caller(0)
	// load template，在配置路由时会初始化一些模板函数，因此放 router 后面
	simplate.SetViewsPath(path.Join(path.Dir(file), "..", "views"))
	simplate.SetLayoutFile("layout/default.html")
	// 添加模板函数
	helper.AddFuncMap()
	simplate.InitTemplate()
	r.HTMLRender = simplate.GinRender

	// 3.2 配置路由
	rR = router.Init(r)

	resultCode := m.Run()
	fmt.Println("end test of controllers")
	os.Exit(resultCode)
}

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	rR.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
