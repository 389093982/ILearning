package models

import (
	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id          		int		`json:"id"`
	CourseName    		string	`json:"course_name"`		// 课程名称
	CourseAuthor		string	`json:"course_author"`		// 课程作者
	CourseType			string	`json:"course_type"`		// 课程内容类型
	CourseShortDes		string	`json:"course_short_desc"`	// 课程简介
	SmallImage			string	`json:"small_image"`		// 课程小图标
	Score				float32	`json:"score"`				// 课程得分
	CourseNumber		int		`json:"course_number"`		// 课程集数
	CourseStatus		string	`json:"course_status"`		// 课程更新状态
	MediaType			string	`json:"media_type"` 		// 课程媒体类型
}

type CourseVedio struct {
	Id					int
	Course				*Course	`orm:"rel(fk)"`
	VedioName			string	// 视频名称
	CourseName			string	// 视频对应的课程名称
	VedioNumber			int		// 视频集数编号
	FirstPlay			string	// 第一存储/播放位置
	SecondPlay			string  // 第二存储/播放位置
}

func QueryCourseExist(course_name string) (count int64, err error)  {
	o := orm.NewOrm()
	count, err = o.QueryTable("course").Filter("course_name", course_name).Count()
	return
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

