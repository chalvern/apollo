package router

import (
	"github.com/chalvern/apollo/app/helper"
	"github.com/chalvern/apollo/app/interceptors"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// Config 路由配置
type Config struct {
	Name         string
	Method       string
	AbsolutePath string
	Handlers     []gin.HandlerFunc
}

type routerConfigs []Config
type routerConfigsMap map[string]Config

func (rcm routerConfigsMap) GetAbsoluteURLOf(name string) string {
	config, ok := rcm[name]
	if !ok {
		return "/"
	}
	return config.AbsolutePath
}

var (
	routerConfigSlice = make([]Config, 0, 20)
	routerConfigMap   = make(routerConfigsMap)
)

// get method
func get(name, absolutePath string, handlers ...gin.HandlerFunc) {
	h := []gin.HandlerFunc{interceptors.JwtMiddleware()}
	h = append(h, handlers...)
	h = append(h, interceptors.AfterRouterMiddleware())
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "get",
		AbsolutePath: absolutePath,
		Handlers:     h,
	})
}

// post method
func post(name, absolutePath string, handlers ...gin.HandlerFunc) {
	h := []gin.HandlerFunc{interceptors.JwtMiddleware()}
	h = append(h, handlers...)
	h = append(h, interceptors.AfterRouterMiddleware())
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "post",
		AbsolutePath: absolutePath,
		Handlers:     h,
	})
}

// put method
func put(name, absolutePath string, handlers ...gin.HandlerFunc) {
	h := []gin.HandlerFunc{interceptors.JwtMiddleware()}
	h = append(h, handlers...)
	h = append(h, interceptors.AfterRouterMiddleware())
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "put",
		AbsolutePath: absolutePath,
		Handlers:     h,
	})
}

// Init initialize engine
func Init(r *gin.Engine) *gin.Engine {

	routerInit()
	adminRouterInit()

	// 把 routerConfigMap 反向注入到 helper
	helper.SetRouterConfig(routerConfigMap)

	for _, rc := range routerConfigSlice {
		routerConfigMap[rc.Name] = rc
		switch m := rc.Method; m {
		case "get":
			r.GET(rc.AbsolutePath, rc.Handlers...)
		case "post":
			r.POST(rc.AbsolutePath, rc.Handlers...)
		case "put":
			r.PUT(rc.AbsolutePath, rc.Handlers...)
		default:
			sugar.Fatalf("方法未注册: %v", rc)
		}
	}

	return r
}
