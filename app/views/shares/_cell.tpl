<div class="media">
	<div class="media-left">
		<div class="cell-a-avatar">
			<a href="/user/info/{{.User.ID}}" style="color:white;">
				<div class="avatar">{{.User.Email | firstChar}}</div>
			</a>
		</div>
	</div>
	<div class="media-body">
		<div class="title">
			<a href="/share/get/{{.ID}}?url={{.URL}}" target="_blank">《{{.Title}}》</a>
		</div>
		<div class="reviews">
			评述：{{.Review}}
		</div>
		<p class="gray">
			推荐自 <a href="/user/info/{{.User.ID}}">{{.User.Email}}</a> · 
			{{ if .Tag}}
			<a href="/tag/s/{{.Tag}}" class="index-share-tag">{{.Tag}}</a> · 
			{{ end }}
		</p>
	</div>
</div>