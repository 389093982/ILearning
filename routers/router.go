package routers

import (
	"ILearning/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{},"get,post:Index")
	beego.Router("/user/regist",&controllers.UserController{},"get,post:Regist")
	beego.Router("/user/login",&controllers.UserController{},"get,post:Login")
	beego.Router("/course/index",&controllers.CourseController{},"get,post:Index")
	beego.Router("/course/queryCourse",&controllers.CourseController{},"get,post:QueryCourse")
	beego.Router("/course/play",&controllers.CourseController{},"get,post:Play")
}
