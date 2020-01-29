{{if .SideTags}}
	<div class="body-sidebar hidden-xs hidden-sm">
		<h4>推荐标签(<a href="{{link `tag_list`}}">查看所有</a>)</h4>
		{{ range .SideTags }}
			<a href="{{link `tag_detail` `t` .Name}}" class="btn btn-default share-tag">{{.Name}}</a>
		{{ end }}
	</div>
{{end}}