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
			<a href="{{link `share_detail` `id` .ID}}" target="_blank">《{{.Title}}》</a>
		</div>
		<div class="reviews">
			评述：{{.Review}}
		</div>
		<p class="gray">
			推荐自 <a href="{{link `user_detail` `uid` .User.ID}}">{{.User.Nickname}}</a> · 
			{{ if .Tag}}
			<a href="{{link `tag_detail` `t` .Tag}}" class="index-share-tag">{{.Tag}}</a>
			{{ end }}
		</p>
	</div>
</div>
<div class="divide mar-top-5"></div>