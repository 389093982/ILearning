package models

import (
	"github.com/astaxie/beego/orm"
)

type Course struct {
	Id          		int		`json:"id"`
	CourseName    		string	`json:"course_name"`		// 课程名称
	CourseAuthor		string	`json:"course_author"`		// 课程作者
	CourseType			string	`json:"course_type"`		// 课程内容类别
	CourseSubType		string 	`json:"course_sub_type"`	// 课程内容子类别
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
	VedioNumber			int		// 视频集数编号
	FirstPlay			string	// 第一存储/播放位置
	SecondPlay			string  // 第二存储/播放位置
}

// 四个参数：课程 id,视频集数,存储文件路径,原始文件名
func UploadVedio(course_id int, vedio_number int, saveFilePath string, fileName string)  (flag bool) {	// 默认 flag 为 false
	o := orm.NewOrm()
	// 查询 course 相关信息
	var course Course
	o.QueryTable("course").Filter("id",course_id).Limit(1).One(&course)
	count, err1 := o.QueryTable("course_vedio").Filter("course_id", course_id).Filter("vedio_number",vedio_number).Count()
	if err1 == nil {
		if count > 0{
			_, err2 := o.QueryTable("course_vedio").Filter("course_id", course_id).Filter("vedio_number",vedio_number).Update(orm.Params{
				"vedio_name": fileName,"first_play": saveFilePath,"second_play": saveFilePath,
			})
			if err2 == nil{
				flag = true
			}
		}else{
			var cVedio = CourseVedio{Course:&course, VedioName:fileName, VedioNumber:vedio_number, FirstPlay:saveFilePath, SecondPlay:saveFilePath}
			_,err3 := o.Insert(&cVedio)
			if err3 == nil{
				flag = true
			}
		}
	}
	if flag{
		// 更新完视频需要更新课程表里面的视频集数
		if vedio_number > course.CourseNumber{
			course.CourseNumber = vedio_number
			if _, err := o.Update(&course); err == nil{
				flag = true
			}
		}
	}
	return
}

func ChangeImage(id int,saveFilePath string) (flag bool){
	o := orm.NewOrm()
	_, err := o.QueryTable("course").Filter("id", id).Update(orm.Params{
		"small_image": saveFilePath,
	})
	if err == nil {
		flag = true
	}else{
		flag = false
	}
	return
}

func AddNewCourse(course *Course) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(course)
	return  id, err
}

func QueryCourseExist(course_name string) (count int64, err error)  {
	o := orm.NewOrm()
	count, err = o.QueryTable("course").Filter("course_name", course_name).Count()
	return
}

func QueryCourseById(id int) (course Course, err error)  {
	o := orm.NewOrm()
	err = o.QueryTable("course").Filter("id", id).Limit(1).One(&course)
	return
}

func QueryCourse(condArr map[string]string, page int, offset int) (courses []Course, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("course")
	cond := orm.NewCondition()

	if _,ok:=condArr["CourseAuthor"];ok{
		cond = cond.And("CourseAuthor", condArr["CourseAuthor"])
		//.Filter("username",user.Username).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
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

