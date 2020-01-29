<div class="row">
	<div class="col-md-9 mt-3"> 
		<div class="panel panel-default">
			<div class="panel-heading index-panel-heading">
				<ul class="nav nav-pills">
					<li id="tab_0">
						<a href="/tags" style="padding: 1px 15px;">标签</a>
					</li>
				</ul>
			</div>
			<div class="panel-body paginate-bot">
				{{range .Tags}}
				    <a href="{{link `tag_detail` `t` .Name}}" class="btn btn-default share-tag">{{.Name}}</a>
				{{end}}
			</div>
			</div>
	</div>
	<div class="col-md-3">
		{{template "home/_sidebar.tpl" . }}
	</div>
</div>