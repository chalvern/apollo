package helper

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2html(t *testing.T) {
	rawStr := "<p>I am paragraph</p>"
	str := Str2html(rawStr)
	assert.IsType(t, template.HTML(""), str)
}

func TestNoHTML(t *testing.T) {
	rawContent := `<script type="text/javascript">var i=0;</script>`
	content := NoHTML(rawContent)
	assert.Contains(t, content, "&lt;script")
	assert.Contains(t, content, "script&gt;")
}

func TestMarkdownHelper(t *testing.T) {
	content := `
## 我的大学

* 高尔基的自传三部曲之一
	`
	str := MarkdownHelper(content)
	assert.NotNil(t, str)
	assert.Contains(t, str, "<h2>我的大学</h2>")
}
