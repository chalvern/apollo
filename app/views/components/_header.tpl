<nav class="navbar">
  <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
                aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="/">{{ brand_title }}</a>
      </div>
      <div id="navbar" class="navbar-collapse collapse header-navbar">
        <ul class="nav navbar-nav navbar-right">
          <li>
            <a href="{{link `home_page`}}">首页</a>
          </li>
          <li>
            <a href="{{link `about`}}">关于</a>
          </li>
          {{if .Account}}
            <li>
              <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown"
               data-hover="dropdown">
                设置<span class="caret"></span>
              </a>
              <span class="dropdown-arrow"></span>
              <ul class="dropdown-menu">
                <li><a href="{{link `user_info`}}">个人资料</a></li>
                <li><a href="{{link `signout`}}">退出</a></li>
              </ul>
            </li>
           {{ else }}
            <li><a href="{{link `signin`}}">登录</a></li>
            <li><a href="{{link `signup`}}">注册</a></li>
          {{ end }}
        </ul>
      </div>
  </div>
</nav>