{{if account_normal_authority .Account}}
<div class="">
    <a href="{{link `share_new_get`}}" class="btn btn-default btn-block apollo-new-share">
        创建新分享
    </a>
</div>
{{ end }}

<div class="body-sidebar">
  {{template "components/_sidebar-about.tpl" . }}
</div>

<div class="body-sidebar">
    {{template "components/_friend_link.tpl" . }}
</div>

{{template "tags/_taglist.tpl" . }}

<div class="body-sidebar">
    <div style="text-align: center;">
        <img src="/static/img/qrcode_jzb.jpg" style="width: 100%;"></img>
        <div>公众号（见周边）</div>
    </div>
</div>