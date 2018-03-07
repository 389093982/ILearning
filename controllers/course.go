package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"ILearning/models"
	"encoding/json"
	"fmt"
	"ILearning/ileaning/util"
)

type CourseController struct {
	beego.Controller
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
	fmt.Print(data)
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}
