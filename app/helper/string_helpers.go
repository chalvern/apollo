package helper

import (
	"html/template"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// NoHTML 对 HTML 标签转义
func NoHTML(str string) string {
	strScriptLeft := strings.Replace(str, "<script", "&lt;script", -1)
	strScriptRight := strings.Replace(strScriptLeft, "script>", "script&gt;", -1)
	return strings.Replace(strScriptRight, "\r\n", "\n", -1)
}

// StringLimitLengthHelper 返回特定长度的内容
func StringLimitLengthHelper(rawStr string, limitLen int) string {
	rawRune := []rune(rawStr)
	if len(rawRune) > limitLen {
		return string(rawRune[:limitLen]) + "..."
	}
	return rawStr
}

// MarkdownHelper markdown转换
func MarkdownHelper(content string) string {
	rawBytes := []byte(NoHTML(content))
	markdownBytes := blackfriday.Run(
		rawBytes,
		// blackfriday.WithExtensions(blackfriday.HardLineBreak),
	)
	return string(markdownBytes)
}

// Str2html Convert string to template.HTML type.
func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}
