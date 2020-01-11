package router

import (
	"fmt"

	"github.com/chalvern/simplate"
	"github.com/chalvern/sugar"
)

// URLPathOfHelper 根据 URL 名称转 path 地址
// params 按照 name=value 的方式分别传入 name 和 value
// 且必须成对出现
func URLPathOfHelper(name string, params ...string) string {
	config, ok := routerConfigMap[name]
	if !ok {
		return "/"
	}

	urlPath := config.AbsolutePath
	// 目前只接受 2 对参数
	if paramsLen := len(params); paramsLen == 2 {
		urlPath = fmt.Sprintf("%s?%s=%s", urlPath, params[0], params[1])
	} else if paramsLen == 4 {
		urlPath = fmt.Sprintf("%s?%s=%s&%s=%s", urlPath, params[0], params[1], params[2], params[3])
	}
	return urlPath
}

func simplateFuncRegistor() {
	sugar.Debug("simplateFuncRegistor")
	simplate.AddFuncMap("URLPathof", URLPathOfHelper)
	simplate.AddFuncMap("link", URLPathOfHelper)
}
