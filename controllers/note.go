package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"ILearning/models"
	"time"
	"github.com/astaxie/beego/orm"
)

type NoteController struct {
	beego.Controller
}

func (this *NoteController) QueryNoteExist()  {
	//初始化
	note_name := this.GetString("note_name")
	user_name := this.Ctx.Input.Session("UserName").(string)

	if strings.TrimSpace(note_name) != ""{
		count, err := models.QueryNoteExist(note_name, user_name)
		if err == nil && count == 0{
			this.Data["json"] = &map[string]interface{}{"flag": false}
		}else{
			this.Data["json"] = &map[string]interface{}{"flag": true}
		}
	}else{
		this.Data["json"] = &map[string]interface{}{"flag": false}
	}
	this.ServeJSON()
}

func (this *NoteController) AddNote() {
	method := this.Ctx.Request.Method
	if method == "GET"{
		this.TplName = "note/note_add.html"
	}else {
		user_name := this.Ctx.Input.Session("UserName").(string)

		var note models.Note
		note.NoteName = this.GetString("note_name")
		note.NoteOwner = user_name
		note.NoteKeyWords = this.GetString("note_key_words")
		note.NoteContent = orm.TextField(this.GetString("note_content"))
		note.CreatedBy = user_name
		note.CreatedTime = time.Now()
		note.LastUpdatedBy = user_name
		note.LastUpdatedTime = time.Now()
		_, err := models.AddNote(&note)
		if err == nil {
			this.Data["json"] = &map[string]interface{}{"flag": true}
		} else {
			this.Data["json"] = &map[string]interface{}{"flag": false, "msg":err.Error()}
		}
		this.ServeJSON()
	}
}

func (this *NoteController) ListNote()  {

}

