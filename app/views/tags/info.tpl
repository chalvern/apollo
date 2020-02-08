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
			<a href="{{link `home_page` `t` `0`}}" style="padding: 1px 15px;">全部</a> | 
      标签：{{ .CurrentTagName }}
      {{if account_manager_authority .Account}}
        | <a href="{{link `tag_edit_get` `t` .CurrentTagName}}">编辑标签</a>
      {{end}}
    </div>
    <div class="panel-body paginate-bot">
      {{ if .CurrentTag }}
        <div class="reviews" style="">
          {{if eq .CurrentPage 1}}
            <!-- 第一页展示全文，其他页展示前 100 个字符 -->
            {{str2html ((str_limit_length .CurrentTag.Desc 2000)|markdown)}}
          {{else}}
            {{str2html ((str_limit_length .CurrentTag.Desc 100)|markdown)}}
          {{end}}
        </div>
      {{end}}
    </div>
  </div>
  
  <div class="status pull-right">
    第 {{.CurrentPage}}/{{.TotalPage}} 页
  </div>
  <div class="panel panel-default">
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
        window.location.href = window.location.pathname + "/?page=" + page + "&t={{.CurrentTagName}}";
      }
    });
  });
</script>