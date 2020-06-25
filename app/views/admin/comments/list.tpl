<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<ul class="nav nav-pills">
				<div>用户展示审核用户提交的评论</div>
			</ul>
		</div>
		<div class="panel-body paginate-bot">
			{{range .Comments}}
				{{template "admin/comments/_cell.tpl" .}}
			{{end}}
			<ul id="page"></ul>
		</div>
	</div>
</div>
<div class="col-md-3">
  {{template "home/_sidebar.tpl" . }}
</div>

</div>

<div class="placeholder-body"></div>