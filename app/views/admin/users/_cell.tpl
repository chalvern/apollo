<div class="media">
	<div class="media-body">
		<div class="title">
			<a href="{{link `user_detail` `uid` .ID}}">{{.Nickname}}</a>
		</div>
		<div class="gray">
			<span>邮箱: {{.Email}}</span> ·
			<span>权限: {{.Priority}}</span>
		</div>
	</div>
</div>
<div class="divide mar-top-5"></div>