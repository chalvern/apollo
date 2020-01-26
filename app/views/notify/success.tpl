
<div id="success-common" style="text-align:center">
</div>
<div class="placeholder-body"></div>

<script type="text/javascript">
    {{ if .Timeout}}
        var t = {{.Timeout}};
    {{ else }}
        var t = 15;
    {{ end }}
    function showTime() {
        {{ if .RedirectName }}
            $("#success-common").html("{{.Info}}，" + t + " 秒后自动跳转到 <a href='{{.RedirectURL}}'>{{.RedirectName}}</a>。");
        {{ else }}
            $("#success-common").html("{{.Info}}，" + t + " 秒后自动跳转到 <a href='/'>首页</a>。");
        {{ end }}
        t -= 1;
        if (t == 0) {
            {{ if .RedirectURL }}
                window.location.href = '{{ .RedirectURL }}';
            {{ else }}
                window.location.href = '/';
            {{ end }}
            return;
        }
        setTimeout("showTime()",1000);
    }
    showTime();
</script>