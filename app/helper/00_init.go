package helper

// RouterConfig 路由配置
type RouterConfig interface {
	GetAbsoluteURLOf(name string) string
}

var (
	routerConfig RouterConfig
)

// SetRouterConfig 设置 routerConfig
func SetRouterConfig(rc RouterConfig) {
	routerConfig = rc
}
