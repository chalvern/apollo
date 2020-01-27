<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">注册</div>
      <div class="alert alert-danger alert-dismissible apollo-alert hide" role="alert">
        <div id="fflash">ji</div>
      </div>
      <div class="panel-body">
        <form action="{{link `signup`}}" id="signup-form" method="post">
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">昵称</label>
            <input type="text" id="nick_name" name="nick_name" class="form-control" placeholder="昵称">
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="email">邮箱</label>
            <input type="text" id="email" onchange="checkEmail(this.value)" name="email" class="form-control" placeholder="邮箱"></input>
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="password">密码（至少 8 位）</label>
            <input type="password" id="password" onchange="checkPassword(this.value)" name="password" class="form-control" placeholder="密码"></input>
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="password2">确认密码</label>
            <input type="password" id="password2" onchange="checkPassword2(this.value)" name="password2" class="form-control" placeholder="确认密码"></input>
          </div>
          <div class="form-group">
            {{template "captcha/captcha.tpl"}}
          </div>
          <input type="submit" class="btn btn-default" onclick="signup();return false;" value="注册"></input> <a href="{{link `signin`}}">去登录</a>
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
// checkEmail 检查邮箱
function checkEmail(email) {
  var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
  if (email.length<4 || !reg.test(email)) {
    $("#fflash").html("请输入正确的邮箱地址");
    $(".apollo-alert").removeClass('hide');
    return;
  } else {
    $(".apollo-alert").addClass('hide');
  }
}

function checkPassword(password) {
  if (password.length < 8) {
    $("#fflash").html("密码至少 8 位");
    $(".apollo-alert").removeClass('hide');
    return false;
  }
  $(".apollo-alert").addClass('hide');
  return true;
}

// checkPassword2 检查两次输入的密码是否一致
function checkPassword2(password2) {
  var password = $("#password").val();
  if (password !== password2) {
    $("#fflash").html("两次输入的密码不一致");
    $(".apollo-alert").removeClass('hide');
    return false;
  }
  $(".apollo-alert").addClass('hide');
  return true;
}

// signup 注册
function signup() {
  var email = $("#email").val();
  var password = $("#password").val();
  var password2 = $("#password2").val();
  var captcha = $("#captcha").val();

  if (!checkPassword(password) || !checkPassword2(password2)) { return false};

  if (!email || !password || !password2 || !captcha) {
    $("#fflash").html("请按照要求填写所有字段内容");
    $(".apollo-alert").removeClass('hide');
    return false;
  }
  $(".apollo-alert").addClass('hide');
  $("#signup-form").submit();
  return true;
}
</script>