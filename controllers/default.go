package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (this *MainController) Index()  {
	coursetypelist := this.GetString("coursetypelist")
	if(coursetypelist == "coursetypelist"){
		this.Data["CourseTypeListShow"] = "CourseTypeListShow"
	}
	this.TplName = "index.html"
}