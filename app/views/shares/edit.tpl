<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">新分享</div>
      <div class="alert alert-danger alert-dismissible jzb-alert hide" role="alert">
        <div id="fflash"></div>
      </div>
      <div class="panel-body">
        <form action="{{link `share_edit_get` `id` .Share.ID}}" id="signin-form" method="post">
          <div class="form-group">
            <em style="color: red;">* </em><label for="url">分享的 url 地址</label>
            <input type="text" id="url" name="url" class="form-control" onchange="analysisURL(this.value)"
                placeholder="url地址,带http(s)，例如 https://jingwei.link" value="{{.Share.URL}}"></input>
          </div>
          <div class="form-group">
            <em style="color: red;">* </em>
            <label for="title">分享标题</label> · 
            <a href="#" onclick="onFresh(); return false;">刷新</a>
            <textarea id="title" name="title" class="form-control" rows="1" placeholder="分享标题">{{.Share.Title}}</textarea>
          </div>
          <div class="form-group">
            <em style="color: red;">* </em><label for="review">简单评论</label>
            <textarea id="review" name="review" class="form-control" rows="3"
                placeholder="陈述了什么事实？说明了什么问题？为什么推荐？">{{.Share.Review}}</textarea>
          </div>
          <div class="form-group">
            <label for="title">标签</label>
            <input type="text" id="tag" name="tag" class="form-control" placeholder="标签，比如：java" value="{{.Share.Tag}}"></input>
          </div>
          <input type="submit" class="btn btn-default" value="提交"></input>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-3">
  {{template "shares/_sidebar.tpl" . }}
  </div>
</div>

<script language=javascript>
// checkEmail 检查邮箱
function analysisURL(url) {
  var regURL = /((([A-Za-z]{3,9}:(?:\/\/)?)(?:[\-;:&=\+\$,\w]+@)?[A-Za-z0-9\.\-]+|(?:www\.|[\-;:&=\+\$,\w]+@)[A-Za-z0-9\.\-]+)((?:\/[\+~%\/\.\w\-_]*)?\??(?:[\-\+=&;%@\.\w_]*)#?(?:[\.\!\/\\\w]*))?)/
  if (!regURL.test(url)) {
    $("#fflash").html("输入的 url 不合法，请检查是否存在问题");
    $(".jzb-alert").removeClass('hide');
    return false;
  }
  $(".jzb-alert").addClass('hide');
  $.ajax({
    url: "{{link `url_title`}}?url=" + url,
    success: function(result){
      $("#title").val(result);
    },
    timeout: 1000 * 5,
    error: function(e) {
      $("#title").val(result);
    },
});
}

function onFresh() {
    analysisURL($("#url").val());
}
</script>