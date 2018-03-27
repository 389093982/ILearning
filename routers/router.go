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
	beego.Router("/user/logout",&controllers.UserController{},"get,post:Logout")
	beego.Router("/course/index",&controllers.CourseController{},"get,post:Index")
	beego.Router("/course/queryCourse",&controllers.CourseController{},"get,post:QueryCourse")
	beego.Router("/course/play",&controllers.CourseController{},"get,post:Play")
	beego.Router("/course/homemanage",&controllers.CourseController{},"get,post:HomeManage")
	beego.Router("/course/newcourse",&controllers.CourseController{},"get,post:NewCourse")
	beego.Router("/course/courselist",&controllers.CourseController{},"get,post:CourseList")
	beego.Router("/course/showCourseDetail",&controllers.CourseController{},"get,post:ShowCourseDetail")
	beego.Router("/course/newcourse/add",&controllers.CourseController{},"get,post:AddNewCourse")
	beego.Router("/course/newcourse/changeImage",&controllers.CourseController{},"get,post:ChangeImage")
	beego.Router("/course/newcourse/uploadvedio",&controllers.CourseController{},"get,post:UploadVedio")
	beego.Router("/course/newcourse/endUpdate",&controllers.CourseController{},"get,post:EndUpdate")
	beego.Router("/course/queryCourseExist",&controllers.CourseController{},"get,post:QueryCourseExist")
	beego.Router("/course/favorite/toggle",&controllers.CourseController{},"get,post:ToggleFavorite")

	beego.Router("/comment/topicTheme/filter",&controllers.CommentController{},"get,post:FilterTopicTheme")
	beego.Router("/comment/topicReply/add",&controllers.CommentController{},"get,post:AddTopicReply")
	beego.Router("/comment/topicReply/filter",&controllers.CommentController{},"get,post:FilterTopicReply")

	beego.Router("/note/list",&controllers.NoteController{},"get,post:ListNote")
	beego.Router("/note/add",&controllers.NoteController{},"get,post:AddNote")
	beego.Router("/note/queryNoteExist",&controllers.NoteController{},"get,post:QueryNoteExist")
}
