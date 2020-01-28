<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">注册</div>
      <div class="alert alert-danger alert-dismissible apollo-alert hide" role="alert">
        <div id="fflash">ji</div>
      </div>
      <div class="panel-body">
        <form action="{{link `admin_account_edit_get` `uid` .User.ID}}" id="signup-form" method="post">
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">昵称:</label> {{.User.Nickname}}
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">邮箱</label> {{.User.Email}}
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">Priority</label>
            <input type="text" id="priority" name="priority" class="form-control" value={{.User.Priority}}>
          </div>
          <input type="submit" class="btn btn-default" value="更新"></input>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-3">
    <div class="body-sidebar hidden-xs hidden-sm">
      {{template "components/_sidebar-about.tpl" . }}
    </div>
  </div>
</div>