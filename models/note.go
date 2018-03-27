package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Note struct {
	Id          			int			`json:"id"`
	NoteName				string		`json:"note_name"`			// 笔记名称
	NoteOwner				string		`json:"note_owner"`			// 笔记作者
	NoteKeyWords			string		`json:"note_key_words"`		// 笔记关键字,用于模糊搜索使用
	NoteContent 			orm.TextField `json:"note_content"`		// 笔记内容
	CreatedBy				string 		`json:"created_by"`			// 创建人
	CreatedTime				time.Time	`json:"created_time"`		// 创建时间
	LastUpdatedBy			string		`json:"last_updated_by"`	// 修改人
	LastUpdatedTime			time.Time	`json:"last_updated_time"`	// 修改时间
}

func QueryNoteExist(note_name, user_name string) (count int64, err error)  {
	o := orm.NewOrm()
	count, err = o.QueryTable("note").Filter("note_name", note_name).Filter("note_owner", user_name).Count()
	return
}

func AddNote(note *Note) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(note)
	return  id, err
}
