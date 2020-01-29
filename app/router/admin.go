package router

import (
	"github.com/chalvern/apollo/app/controllers/admin"
	i "github.com/chalvern/apollo/app/interceptors"
	"github.com/chalvern/apollo/app/model"
)

func adminRouterInit() {
	get("admin_account_list", "/admin/account/list",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsList)
	get("admin_account_edit_get", "/admin/account/edit",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsEditGet)
	post("admin_account_edit_post", "/admin/account/edit",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsEditPost)
}
