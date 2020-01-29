{{if account_normal_authority .Account}}
<div class="">
    <a href="{{link `share_new_get`}}" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}

<div class="body-sidebar hidden-xs hidden-sm">
  {{template "components/_sidebar-about.tpl" . }}
</div>

<div class="body-sidebar hidden-xs hidden-sm">
    {{template "components/_friend_link.tpl" . }}
</div>

{{template "tags/_taglist.tpl" . }}