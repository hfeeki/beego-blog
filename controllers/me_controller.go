package controllers

type MeController struct {
	AdminController
}

func (this *MeController) Default() {
	this.Layout = "layout/admin.html"
	this.TplNames = "me/default.html"
}
