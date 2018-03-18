package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"ILearning/models"
	"encoding/json"
	"fmt"
	"ILearning/ileaning/util"
	"strings"
	"path"
	"github.com/satori/go.uuid"
)

type CourseController struct {
	beego.Controller
}

var UploadFileSavePathImg string

func init(){
	UploadFileSavePathImg = beego.AppConfig.String("UploadFileSavePathImg")
}

func (this *CourseController) UploadVedio()  {
	method := this.Ctx.Request.Method
	if method == "GET"{
		// get 请求加载页面
		this.Layout = "course/home_manage.html"
		this.TplName = "course/upload_vedio.html"
	}else{
		// post 请求更新视频
	}

}

func (this *CourseController) ChangeImage()  {
	id, _:= this.GetInt("id")
	f, fh, err := this.GetFile("file")
	defer f.Close()
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"path": "", "status": "ERROR"}
		this.ServeJSON()
	} else {
		// 与 this.GetFile("file") 保持一致的名字
		// fh.Filename 原始文件名,存储时使用 UUID 进行重命名
		u,_ := uuid.NewV4()
		newFileName := u.String() + path.Ext(fh.Filename)
		saveFilePath := path.Join(UploadFileSavePathImg, newFileName)
		err := this.SaveToFile("file", saveFilePath)
		// 更新图片
		flag := models.ChangeImage(id, "/" + saveFilePath)
		if(err == nil && flag == true){
			this.Data["json"] = &map[string]interface{}{"path": saveFilePath, "status": "SUCCESS"}
		}else{
			this.Data["json"] = &map[string]interface{}{"path": saveFilePath, "status": "ERROR"}
		}
		this.ServeJSON()
	}
}

func (this *CourseController) CourseList()  {
	// 查看发布课程
	this.Layout = "course/home_manage.html"
	this.TplName = "course/course_list.html"
}

func (this *CourseController) AddNewCourse()  {
	//初始化
	data := make(map[string]interface{}, 1)

	var course models.Course
	course_name := this.GetString("course_name")
	course_type := this.GetString("course_type")
	course_sub_type := this.GetString("course_sub_type")
	course_short_desc := this.GetString("course_short_desc")
	course.CourseName = course_name
	course.CourseType = course_type
	course.CourseSubType = course_sub_type
	course.CourseShortDes = course_short_desc
	course.CourseAuthor = this.Ctx.Input.Session("UserName").(string)
	_, err := models.AddNewCourse(&course)
	if err == nil{
		data["status"] = "SUCCESS"
		data["msg"] = "保存成功!"
	}else{
		data["status"] = "ERROR"
		data["msg"] = err.Error()
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

	filterType := this.GetString("filterType","")
	if filterType == "courselist"{
		// filterType == "courselist" 时,查看当前登录用户已发布课程
		condArr["CourseAuthor"] = this.Ctx.Input.Session("UserName").(string)
	}else{
		// 否则从请求参数中获取相关信息
		CourseAuthor := this.GetString("CourseAuthor","")
		CourseType := this.GetString("CourseType","")
		if CourseAuthor != ""{
			condArr["CourseAuthor"] = CourseAuthor
		}
		if CourseType != ""{
			condArr["CourseType"] = CourseType
		}
	}
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
