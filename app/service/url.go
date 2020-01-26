package service

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/chalvern/sugar"
	"github.com/go-resty/resty/v2"
)

// QueryTitleFormURL 提取标题
func QueryTitleFormURL(urlRaw string) (string, error) {
	resp, err := resty.New().R().Get(urlRaw)
	if err != nil || resp.StatusCode() != http.StatusOK {
		sugar.Warnf("resty 获取 %s 出错(statusCode: %d): %v", urlRaw, resp.StatusCode(), err)
		return "", err
	}

	respString := resp.String()
	var regTitle *regexp.Regexp
	if strings.HasPrefix(urlRaw, "https://mp.weixin.qq.com") {
		regTitle, _ = regexp.Compile(`var msg_title = "(.*)";`)
	} else {
		regTitle, _ = regexp.Compile(`<title>(.*)</title>`)
	}

	titles := regTitle.FindStringSubmatch(respString)
	if len(titles) > 0 {
		return titles[1], nil
	}
	return "自动提取失败", nil
}
