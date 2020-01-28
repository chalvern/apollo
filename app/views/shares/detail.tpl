<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<a href="{{link `home_page`}}">首页</a> | 
			<a href="{{link `tag_detail` `t` .Share.Tag}}">{{.Share.Tag}}</a>
		</div>
		<div class="panel-body paginate-bot detail">
			<div class="title">
				《{{ .Share.Title }}》
				<span style="font-size: 14px">
					(<a href="{{link `share_direct_jump` `id` .Share.ID}}" target="_blank">原文</a>)
				</span>
			</div>
			<div class="detail-status">
				作者：<a href="{{link `user_detail` `uid` .Share.User.ID}}">{{.Share.User.Nickname}}</a> · 
				点击：{{.Share.ClickCount}} 次
			</div>
			<div class="divide mar-top-5"></div>
			<div class="reviews">
				评述：{{.Share.Review}}
			</div>
		</div>
	</div>
	<div class="panel panel-default">
		<div class="panel-body paginate-bot">
			{{if .Shares}}
				{{range .Shares}}
					{{template "shares/_cell.tpl" . }}
				{{end}}
			{{ else }}
				<p>暂无评论</p>
			{{ end }}
			<ul id="page"></ul>
		</div>
	</div>
</div>
<div class="col-md-3">
	{{template "home/_sidebar.tpl" . }}
</div>
</div>