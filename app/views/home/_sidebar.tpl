{{if account_normal .Account}}
<div class="">
    <a href="{{link `share_new_get`}}" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}

<div class="body-sidebar hidden-xs hidden-sm">
  {{template "components/_sidebar-about.tpl" . }}
</div>