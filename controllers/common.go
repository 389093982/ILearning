package controllers

import (
	"github.com/astaxie/beego"
	"ILearning/models"
)

type CommonController struct {
	beego.Controller
}

func (this *CommonController) ToggleFavorite()  {
	// 获取课程 id
	favorite_id, _ := this.GetInt("favorite_id")
	favorite_type := this.GetString("favorite_type")
	user_name := this.Ctx.Input.Session("UserName").(string)
	flag := models.IsFavorite(user_name,favorite_id,favorite_type)
	if flag{
		models.DelFavorite(user_name,favorite_id,favorite_type)
	}else{
		models.AddFavorite(user_name,favorite_id,favorite_type)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}
