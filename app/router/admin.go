package router

import (
	"github.com/chalvern/apollo/app/controllers/admin"
	i "github.com/chalvern/apollo/app/interceptors"
)

func adminRouterInit() {
	get("admin_account_list", "/admin/account/list", i.AdminUserMiddleware(), admin.AccountsList)
	get("admin_account_edit_get", "/admin/account/edit", i.AdminUserMiddleware(), admin.AccountsEditGet)
	post("admin_account_edit_post", "/admin/account/edit", i.AdminUserMiddleware(), admin.AccountsEditPost)
}
