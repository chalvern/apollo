{{if .Account}}
<div class="hidden-md hidden-lg">
    <a href="/share" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}
<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<ul class="nav nav-pills">
				<li id="tab_1">
					<a href="{{link `home_page` `t` `0`}}" style="padding: 1px 15px;">全部</a>
				</li>
				<li id="tab_0">
					<a href="{{link `home_page` `t` `1`}}" style="padding: 1px 15px;">精选</a>
				</li>
			</ul>
		</div>
		<div class="panel-body paginate-bot">
			{{range .Shares}}
				{{template "shares/_cell.tpl" .}}
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

<script type="text/javascript">
  $(function () {
	$("#tab_{{.TabIndex}}").addClass("active");
    $("#page").bootstrapPaginator({
      currentPage: '{{.CurrentPage}}',
      totalPages: '{{.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        var t = {{.TabIndex}};
        if (t > 0) {
          window.location.href = "/?p=" + page + "&t={{.TabIndex}}"
        } else {
          window.location.href = "/?p=" + page
        }
      }
    });
  });
</script>