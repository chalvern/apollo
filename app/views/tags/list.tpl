<div class="row">
  <div class="col-md-9 mt-3"> 
    <div class="panel panel-default">
      <div class="panel-heading index-panel-heading">
        <a href="{{link `home_page`}}">首页</a> | 
        <a href="{{link `tag_list`}}" style="padding: 1px 15px;">所有标签</a>
        {{if account_manager .Account}}
          | <a href="{{link `tag_new_get`}}">新建标签</a>
        {{end}}
      </div>
      <div class="panel-body paginate-bot">
        {{range .Tags}}
            <a href="{{link `tag_detail` `t` .Name}}" class="btn btn-default share-tag">
              {{.Name}} ({{.Count}})
            </a>
        {{end}}
      </div>
      </div>
  </div>
  <div class="col-md-3">
    {{template "home/_sidebar.tpl" . }}
  </div>
</div>