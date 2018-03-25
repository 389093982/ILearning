package controllers

import (
	"github.com/astaxie/beego"
	"ILearning/models"
)

type NoteController struct {
	beego.Controller
}

func (this *NoteController) ListNote()  {
	// 获取课程 id
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	topic_theme, err:= models.FilterTopicTheme(topic_id, topic_type)

	if err == nil{
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "topic_theme": topic_theme}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

