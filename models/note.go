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

	TopicId    				int			`json:"topic_id"`			// 评论主题 id
	TopicType				string		`json:"topic_type"`			// 评论主题类型
	NoteContent 			orm.TextField `json:"note_content"`		// 笔记内容
	CreatedBy				string 		`json:"created_by"`			// 创建人
	CreatedTime				time.Time	`json:"created_time"`		// 创建时间
	LastUpdatedBy			string		`json:"last_updated_by"`	// 修改人
	LastUpdatedTime			time.Time	`json:"last_updated_time"`	// 修改时间
}

