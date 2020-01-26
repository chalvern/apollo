{{if .Account}}
<div class="">
    <a href="/share" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}

<div class="body-sidebar hidden-xs hidden-sm">
  {{template "components/_sidebar-about.tpl" . }}
</div>