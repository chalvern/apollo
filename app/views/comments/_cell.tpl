<div class="media">
	<div class="media-left">
		<div class="cell-a-avatar">
			<a href="{{link `user_detail` `uid` .User.ID}}" style="color:white;">
				<div class="avatar">{{.User.Nickname | firstChar}}</div>
			</a>
		</div>
	</div>
	<div class="media-body">
		<div class="title">
			<a href="{{link `user_detail` `uid` .User.ID}}">{{.User.Nickname}}</a>
		</div>
		<div class="reviews">
			{{str2html (.Reply|markdown)}}
		</div>
	</div>
	<div class="media-right">#{{.Number}}</div>
</div>
<div class="divide mar-top-5"></div>