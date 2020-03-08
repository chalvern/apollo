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
				经 <a href="{{link `user_detail` `uid` .Share.User.ID}}">{{.Share.User.Nickname}}</a> 推荐 · 
				{{time_internal_desc .Share.UpdatedAt}}更新 · 
				点击 {{.Share.ClickCount}} 次
				{{if account_has_share_edit_authority .Share .Account}}
				 · <a href="{{link `share_edit_get` `id` .Share.ID}}">编辑</a>
				{{end}}
			</div>
			<div class="divide mar-top-5"></div>
			<div class="reviews">
				评述：{{(.Share.Review|markdown)}}
			</div>
		</div>
	</div>
	<div class="panel panel-default">
		<div class="panel-body paginate-bot">
			{{if .Comments}}
				{{range .Comments}}
					{{template "comments/_cell.tpl" . }}
				{{end}}
				<ul id="page"></ul>
			{{ else }}
				<p>暂无评论</p>
			{{ end }}
		</div>
	</div>
	
	<div class="panel panel-default">
		<div class="panel-body">
			{{if account_normal_authority .Account}}
				<form action="{{link `comment_new_post`}}" method="post">
					<input type="hidden" value="{{.Share.ID}}" name="share_id">
					<div class="form-group">
						<em style="color: red;">* </em>
						<label for="replay">回复</label>
						<textarea id="replay" name="replay" class="form-control" rows="3" placeholder="回复"></textarea>
					</div>
					<button type="submit" class="btn btn-default">回复</button>
				</form>
			{{else}}
				<div><a href="{{link `signin`}}">未登陆</a>或尚未认证授权</div>
			{{end}}
		</div>
	</div>
</div>
<div class="col-md-3">
	{{template "home/_sidebar.tpl" . }}
</div>
</div>

{{if not .Comments}}
<div class="placeholder-body"></div>
{{end}}