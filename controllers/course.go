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
	"time"
)

type CourseController struct {
	beego.Controller
}

var UploadFileSavePathImg string
var UploadFileSavePathVedio string

func init(){
	UploadFileSavePathImg = beego.AppConfig.String("UploadFileSavePathImg")
	UploadFileSavePathVedio = beego.AppConfig.String("UploadFileSavePathVedio")
}

func (this *CourseController) ToggleFavorite()  {
	// 获取课程 id
	course_id, _ := this.GetInt("course_id")
	favorite_type := this.GetString("favorite_type")
	user_name := this.Ctx.Input.Session("UserName").(string)
	count, err:= models.QueryFavorite(user_name,course_id,favorite_type)
	if err == nil{
		if count > 0{
			models.DelFavorite(user_name,course_id,favorite_type)
		}else{
			models.AddFavorite(user_name,course_id,favorite_type)
		}
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CourseController) ShowCourseDetail()  {
	// 获取课程 id
	id, _ := this.GetInt("id")
	course, _:= models.QueryCourseById(id)
	cVedios, _:= models.QueryCourseVedio(id)
	user_name := this.Ctx.Input.Session("UserName").(string)
	count, _:= models.QueryFavorite(user_name,id,"collect")
	if count > 0{
		this.Data["Collect"] = true
	}else{
		this.Data["Collect"] = false
	}
	count1, _:= models.QueryFavorite(user_name,id,"praise")
	if count1 > 0{
		this.Data["Praise"] = true
	}else{
		this.Data["Praise"] = false
	}
	this.Data["Course"] = course
	this.Data["CourseVideo"] = cVedios
	// 视频详情页面
	this.TplName = "course_detail.html"
}

func (this *CourseController) EndUpdate() {
	// 获取课程 id
	id, err := this.GetInt("id")
	if err == nil{
		flag := models.EndUpdate(id)
		if(flag == true){
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		}else{
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		}
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CourseController) UploadVedio()  {
	// 获取课程 id
	id, err1 := this.GetInt("id")
	vedio_number, err2 := this.GetInt("vedio_number")
	f, fh, err3 := this.GetFile("uploadVedioFile")
	defer f.Close()
	// 检查文件格式是否是视频格式
	if !util.CheckVedio(path.Ext(fh.Filename)){
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg":"视频格式不合法!"}
		this.ServeJSON()
		return
	}
	if err1 == nil && err2 == nil && err3 == nil {
		// fh.Filename 原始文件名,存储时使用 UUID 进行重命名
		u,_ := uuid.NewV4()
		newFileName := u.String() + path.Ext(fh.Filename)
		saveFilePath := path.Join(UploadFileSavePathVedio, newFileName)
		// 与 this.GetFile("file") 保持一致的名字
		err := this.SaveToFile("uploadVedioFile", saveFilePath)
		if err == nil{
			// 刷新 DB 记录
			flag := models.UploadVedio(id, vedio_number, "/" + saveFilePath, fh.Filename)
			if(flag == true){
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "msg":"保存成功!"}
			}else{
				this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg":"保存失败!"}
			}
		}else{
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg":"保存失败!"}
		}
		this.ServeJSON()
	}else{
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg":"保存失败!"}
		this.ServeJSON()
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
	user_name := this.Ctx.Input.Session("UserName").(string)

	var course models.Course
	course_name := this.GetString("course_name")
	course_type := this.GetString("course_type")
	course_sub_type := this.GetString("course_sub_type")
	course_short_desc := this.GetString("course_short_desc")
	course.CourseName = course_name
	course.CourseType = course_type
	course.CourseSubType = course_sub_type
	course.CourseShortDes = course_short_desc
	course.CourseStatus = "更新中"
	course.CourseAuthor = user_name
	id, err := models.AddNewCourse(&course)
	if err == nil{
		data["status"] = "SUCCESS"
		data["msg"] = "保存成功!"
	}else{
		data["status"] = "ERROR"
		data["msg"] = err.Error()
	}

	topic_theme := models.TopicTheme{}
	topic_theme.TopicId = int(id)
	topic_theme.TopicType = "course_comment"
	topic_theme.TopicContent = strings.Join([]string{user_name,"@",course_name,"@",course_short_desc}, "")
	topic_theme.CreatedBy = user_name
	topic_theme.CreatedTime = time.Now()
	topic_theme.LastUpdatedBy = user_name
	topic_theme.LastUpdatedTime = time.Now()
	// 增加一条评论主题
	models.AddTopicTheme(&topic_theme)

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
	course_id,_ := this.GetInt("course_id")
	vedio_id,_ := this.GetInt("vedio_id")
	cVedios, _ := models.QueryCourseVedio(course_id)
	for _,value := range cVedios{
		cVedio := &value
		if cVedio.VedioNumber == vedio_id{
			this.Data["CourseVedio"] = &cVedio
			break
		}
	}
	// 播放次数加 1
	models.UpdateWatchNumber(course_id)
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
