package router

import (
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

var (
	routerConfigSlice = make([]Config, 1, 20)
	routerConfigMap   = make(map[string]Config)
)

// get method
func get(name, absolutePath string, handlers ...gin.HandlerFunc) {
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "get",
		AbsolutePath: absolutePath,
		Handlers:     handlers,
	})
}

// post method
func post(name, absolutePath string, handlers ...gin.HandlerFunc) {
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "post",
		AbsolutePath: absolutePath,
		Handlers:     handlers,
	})
}

// put method
func put(name, absolutePath string, handlers ...gin.HandlerFunc) {
	routerConfigSlice = append(routerConfigSlice, Config{
		Name:         name,
		Method:       "put",
		AbsolutePath: absolutePath,
		Handlers:     handlers,
	})
}

// Init initialize engine
func Init(r *gin.Engine) *gin.Engine {

	router()

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
			panic("方法未注册")
		}
	}
	return r
}
