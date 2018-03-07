package models

import (
	"github.com/astaxie/beego/orm"
)

type CourseVedio struct {
	Id					int
	Course
	first_play			string	`第一播放位置`
	second_play			string  `第二播放位置`
}

type Course struct {
	Id          		int
	CourseName    		string	`课程名称`
	CourseAuthor		string	`课程作者`
	CourseType			string	`课程内容类型`
	CourseShortDes		string	`课程简介`
	SmallImage			string	`课程小图标`
	Score				float32	`课程得分`
	CourseNumber		int		`课程集数`
	CourseStatus		string	`课程更新状态`
	MediaType			string 	`课程媒体类型`
}

func QueryCourse(condArr map[string]string, page int, offset int) (courses []Course, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("course")
	cond := orm.NewCondition()

	if condArr["title"] != "" {
		cond = cond.And("title__icontains", condArr["title"])
		//.Filter("username",user.Username).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
	}
	if condArr["keywords"] != "" {
		cond = cond.Or("keywords__icontains", condArr["keywords"])
	}

	qs = qs.SetCond(cond)
	counts,_ = qs.Count()

	qs = qs.Limit(offset, (page - 1) * offset)
	qs.All(&courses)
	//for _, v := range querysOrder {
	//	qs = qs.OrderBy(v)
	//}
	qs.All(&courses)
	return
}

