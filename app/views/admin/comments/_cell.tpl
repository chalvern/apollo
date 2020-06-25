<div class="media">
	<div class="media-body">
		<div class="title">
			<a href="{{link `share_detail` `id` .ShareID}}">{{.Reply}}</a>
		</div>
		<div class="gray">
			<span>{{year_date_str .CreatedAt}} </span>·
			<a href="{{link `user_detail` `uid` .User.ID}}">{{.User.Nickname}}</a>·
			<a href="#">编辑</a>
		</div>
	</div>
</div>
<div class="divide mar-top-5"></div>