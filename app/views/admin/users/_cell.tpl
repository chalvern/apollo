<div class="media">
	<div class="media-body">
		<div class="title">
			<a href="{{link `user_detail` `uid` .ID}}">{{.Nickname}}</a>
		</div>
		<div class="gray">
			<span>邮箱: {{.Email}}</span> ·
			<span>权限: {{.Priority}}</span>
			| <a href="{{link `admin_account_edit_get` `uid` .ID}}">编辑</a>
		</div>
	</div>
</div>
<div class="divide mar-top-5"></div>