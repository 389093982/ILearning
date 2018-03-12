package controllers

import (
	"github.com/astaxie/beego"
	"ILearning/models"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Logout()  {
	this.DelSession("username")
	this.Redirect("/user/login", 302)
}

func (this *UserController) Regist()  {
	Method := this.Ctx.Request.Method
	if Method == "GET"{
		this.TplName = "regist.html"
	}else{
		var user models.User
		inputs := this.Input()
		user.UserName = inputs.Get("username")
		user.PassWd = inputs.Get("passwd")

		err := models.SaveUser(user)

		//初始化
		data := make(map[string]interface{}, 1)

		if err == nil{
			data["status"] = "SUCCESS"
		}else{
			data["Status"] = "ERROR"
			data["ErrorCode"] = err.Error()
			data["ErrorMsg"] = err.Error()
		}
		//序列化
		json_obj, err := json.Marshal(data)
		if err == nil {
			this.Data["json"] = string(json_obj)
		}
		this.ServeJSON()
	}
}

func (this *UserController) Login()  {
	Method := this.Ctx.Request.Method
	if Method == "GET"{
		this.TplName = "login.html"
	}else{
		inputs := this.Input()
		username := inputs.Get("username")
		passwd := inputs.Get("passwd")

		user, err := models.QueryUser(username,passwd)

		//初始化
		data := make(map[string]interface{}, 1)

		if err == nil && &user != nil{
			data["Status"] = "SUCCESS"
			// 将用户名添加到 session 中去
			this.SetSession("UserName",user.UserName)
		}else{
			data["Status"] = "ERROR"
			data["ErrorCode"] = err.Error()
			data["ErrorMsg"] = err.Error()
		}
		//序列化
		json_obj, err := json.Marshal(data)
		if err == nil {
			this.Data["json"] = string(json_obj)
		}
		this.ServeJSON()
	}
}