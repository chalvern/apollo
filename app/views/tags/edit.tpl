<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">更新标签</div>
      <div class="alert alert-danger alert-dismissible jzb-alert hide" role="alert">
        <div id="fflash"></div>
      </div>
      <div class="panel-body">
        <form action="{{link `tag_edit_get` `t` .Tag.Name}}" id="signin-form" method="post">
          <div class="form-group">
            <em style="color: red;">* </em><label for="name">标签名</label>
            <input type="text" id="name" name="name" class="form-control"
                placeholder="标签名（小写）" value="{{.Tag.Name}}"></input>
          </div>
          <div class="form-group">
            <label for="hierarchy">层级（默认为0）</label>
            <input type="text" id="hierarchy" name="hierarchy" class="form-control" 
							placeholder="层级" value="{{.Tag.Hierarchy}}"></input>
          </div>
          <div class="form-group">
            <label for="parent">父标签</label>
            <input type="text" id="parent" name="parent" class="form-control"
                placeholder="父标签" value="{{.Tag.Parent}}"></input>
          </div>
          <div class="form-group">
            <label for="desc">描述</label>
            <textarea id="desc" name="desc" class="form-control" rows="10"
                placeholder="描述">{{.Tag.Desc}}</textarea>
          </div>
          <input type="submit" class="btn btn-default" value="提交"></input>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-3">
  {{template "home/_sidebar.tpl" . }}
  </div>
</div>