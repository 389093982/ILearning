package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"ILearning/models"
	"encoding/json"
	"fmt"
	"ILearning/ileaning/util"
	"strings"
)

type CourseController struct {
	beego.Controller
}

func (this *CourseController) QueryCourseExist()  {
	//初始化
	data := make(map[string]interface{}, 1)
	course_name := this.GetString("course_name","")
	if strings.TrimSpace(course_name) != ""{
		count, err := models.QueryCourseExist(course_name)
		if err == nil && count == 0{
			data["flag"] = false
		}else{
			data["flag"] = true
		}
	}else{
		data["flag"] = true
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}

func (this *CourseController) NewCourse()  {
	// 课程管理界面
	this.Layout = "course/home_manage.html"
	this.TplName = "course/new_course.html"
}

func (this *CourseController) HomeManage()  {
	// 课程管理界面
	this.Layout = "course/home_manage.html"
	this.TplName = "course/home_manage_default.html"
}

func (this *CourseController) Play() {
	// 视频播放
	this.TplName = "course_play.html"
}

func (this *CourseController) Index() {
	this.Data["ExpandExcCrse"] = this.GetString("expandExcCrse", "false")
	this.TplName = "course.html"
}

func (this *CourseController) QueryCourse() {
	condArr := make(map[string]string)
	offset,_ := this.GetInt("offset", 10)	// 每页记录数
	condArr["CourseAuthor"] = this.GetString("CourseAuthor")
	condArr["CourseType"] = this.GetString("CourseType")


	courses, count, err := models.QueryCourse(condArr, 1, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["courses"] = courses
		data["paginator"] = util.Paginator(paginator.Page(),paginator.PerPageNums,paginator.Nums())
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}
