<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">登录</div>
      <div class="alert alert-danger alert-dismissible jzb-alert hide" role="alert">
        <div id="fflash"></div>
      </div>
      <div class="panel-body">
        <form action="/signin" id="signin-form" method="post">
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">邮箱</label>
            <input type="text" id="email" name="email" class="form-control" placeholder="邮箱">
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="password">密码（至少 8 位）</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="密码">
          </div>
          <div class="form-group">
            {{template "captcha/captcha.tpl"}}
          </div>
          <input type="submit" class="btn btn-default" onclick="signin();return false;" value="登录"> <a href="/signup">去注册</a>
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

<script language=javascript>
function signin() {
  var email = $("#email").val();
  var password = $("#password").val();
  var captcha = $("#captcha").val();
  if (!email || !password || !captcha) {
    $("#fflash").html("请按照要求填写所有字段内容");
    $(".jzb-alert").removeClass('hide');
    return false;
  }
  $(".jzb-alert").addClass('hide');
  $("#signin-form").submit();
  return true;
}
</script>