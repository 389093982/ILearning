package models

import (
	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id          		int
	CourseName    		string
	CourseAuthor		string
	CourseType			string
	CourseShortDes		string
	SmallImage			string
	Score				float32
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

