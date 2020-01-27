<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<a href="{{link `home_page`}}">首页</a> | 
			<a href="{{link `tag_info` `t` .Share.Tag}}">{{.Share.Tag}}</a>
		</div>
		<div class="panel-body paginate-bot">
			<div class="title">
				{{ .Share.Title }} · （<a href="{{.Share.URL}}">跳转原文</a>）
			</div>
			<div class="reviews">
				评述：{{.Share.Review}}
			</div>
		</div>
		<div class="panel-body paginate-bot">
			{{range .Shares}}
				{{template "shares/_cell.tpl" . }}
			{{end}}
			<ul id="page"></ul>
		</div>
	</div>
</div>
<div class="col-md-3">
	{{template "home/_sidebar.tpl" . }}
</div>
</div>