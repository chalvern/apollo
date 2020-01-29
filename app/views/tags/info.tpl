{{if .Account}}
<div class="hidden-md hidden-lg">
    <a href="/share" class="btn btn-default btn-block jzb-new-share">
        创建新分享
    </a>
</div>
{{ end }}
<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			标签：{{ .CurrentTag }}
      {{if account_manager_authority .Account}}
        | <a href="{{link `tag_edit_get` `t` .CurrentTag}}">编辑标签</a>
      {{end}}
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

<script type="text/javascript">
  $(function () {
    $("#page").bootstrapPaginator({
      currentPage: '{{.CurrentPage}}',
      totalPages: '{{.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = window.location.pathname + "/?page=" + page + "&uid={{.User.ID}}";
      }
    });
  });
</script>