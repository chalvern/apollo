<div class="media">
	<div class="media-left">
		<div class="cell-a-avatar">
			<a href="{{link `user_info` `uid` .User.ID}}" style="color:white;">
				<div class="avatar">{{.User.NickName | firstChar}}</div>
			</a>
		</div>
	</div>
	<div class="media-body">
		<div class="title">
			<a href="{{link `ping_pong`}}" target="_blank">《{{.Title}}》</a>
		</div>
		<div class="reviews">
			评述：{{.Review}}
		</div>
		<p class="gray">
			推荐自 <a href="{{link `user_info` `uid` .User.ID}}">{{.User.NickName}}</a> · 
			{{ if .Tag}}
			<a href="{{link `ping_pong`}}" class="index-share-tag">{{.Tag}}</a>
			{{ end }}
		</p>
	</div>
</div>
<div class="divide mar-top-5"></div>