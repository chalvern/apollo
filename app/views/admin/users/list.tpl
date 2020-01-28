<div class="row">
<div class="col-md-9 mt-3"> 
	<div class="panel panel-default">
		<div class="panel-heading index-panel-heading">
			<ul class="nav nav-pills">
				<div>权力越大，责任越大！</div>
			</ul>
		</div>
		<div class="panel-body paginate-bot">
			{{range .Users}}
				{{template "admin/users/_cell.tpl" .}}
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
    $("#page").bootstrapPaginator({
      currentPage: '{{.CurrentPage}}',
      totalPages: '{{.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = "/?page=" + page
      }
    });
  });
</script>