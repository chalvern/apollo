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
				{{if .Share.URL}}
					<span style="font-size: 14px">
						(<a href="{{link `share_direct_jump` `id` .Share.ID}}" target="_blank">原文</a>)
					</span>
				{{end}}
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
				{{str2html (.Share.Review|markdown)}}
			</div>
		</div>
	</div>
	{{if .Checklists}}
		<div class="panel panel-default">
			<div class="panel-body paginate-bot">
				<div class="detail-checklist">
					{{range .Checklists}}
						<div class="dropdown">
							<input type="checkbox" class=""> {{.Title}} 
							<a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown" data-hover="dropdown">
								<span class="caret"></span>
							</a>
							<span class="dropdown-arrow"></span>
							<ul class="dropdown-menu">
								<li><a href="javascript:insertNewChecklist({{.ID}});">插入检查项</a></li>
							</ul>
						</div>
					{{end}}
				</div>
			</div>
		</div>
	{{else}}
		<div style="text-align: center;margin-bottom: 20px;">
			<a href="javascript:insertNewChecklist(0);">添加检查项</a>
		</div>
	{{end}}
	{{if account_normal_authority .Account}}
		<div id="checklist-form" class="panel panel-default" style="display:none">
			<div class="panel-body">
				<form action="{{link `checklist_new_post`}}" method="post">
					<input type="hidden" value="{{.Share.ID}}" name="share_id">
					<input type="hidden" id="pre_id" value="0" name="pre_id">
					<div class="form-group">
						<em style="color: red;">* </em>
						<label for="title">添加检查项</label>
						<textarea id="title" name="title" class="form-control" rows="3" placeholder="检查项(不多于200字)"></textarea>
					</div>
					<button type="submit" class="btn btn-default">添加</button>
					<button type="submit" class="btn btn-default" onclick="checklistFormCancel();return false;">关闭</button>
				</form>
			</div>
		</div>
	{{else}}
		<div id="checklist-form">
			<a href="{{link `signin`}}">未登陆</a>或尚未认证授权
		</div>
	{{end}}

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


<script language=javascript>
function checklistFormCancel() {
    document.getElementById("checklist-form").style.display="none";
}

function insertNewChecklist(checklist_id) {
	document.getElementById("pre_id").value = checklist_id;
	document.getElementById("checklist-form").style.display="";
}
</script>