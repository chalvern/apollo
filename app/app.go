package app

import (
	"context"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/chalvern/apollo/app/helper"
	"github.com/chalvern/apollo/app/router"
	"github.com/chalvern/apollo/configs/constants"
	"github.com/chalvern/simplate"
	"github.com/chalvern/sugar"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Run app run
func Run(ctx context.Context) {
	var (
		addr = viper.GetString(constants.WebServerAddr)
		env  = viper.GetString(constants.CoreEnv)
	)

	// 1. production mode for default
	// 1. 默认为生产模式
	if env != constants.EnvDevelop {
		gin.SetMode(gin.ReleaseMode)
	}

	// 2. initial gin
	// 2. 初始化 gin
	r := gin.New()
	r.Use(gin.Recovery())

	// 2.1 设置 logger 为 zap
	// 参考：https://github.com/gin-contrib/zap
	r.Use(ginzap.Ginzap(sugar.GetZapLogger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(sugar.GetZapLogger(), true))

	// 2.2 load template
	_, file, _, _ := runtime.Caller(0)
	simplate.SetViewsPath(path.Join(path.Dir(file), "views"))
	simplate.SetLayoutFile("layout/default.html")
	// 2.2.1 添加模板函数
	helper.AddFuncMap()
	simplate.InitTemplate()
	r.HTMLRender = simplate.GinRender

	// 2.3 static source
	r.Static("/static", path.Join(path.Dir(file), "..", "assets"))

	// 3. config router
	// 3. 配置路由
	rR := router.Init(r)

	// 4. run server
	// 4. 启动 server
	srv := &http.Server{
		Addr:    addr,
		Handler: rR,
	}

	go func() {
		// service connections
		sugar.Infof("Web Server starting: http://%s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalf("listen: %s\n", err)
		}
	}()

	// stop gracefully
	// 5 优雅停止
	<-ctx.Done()

	sugar.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Errorf("Web Server Shutdown Error: %v", err)
	}
	sugar.Info("Web Server exiting")

}
