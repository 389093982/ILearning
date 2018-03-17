package filter

import (
	"github.com/astaxie/beego/context"
	"sync"
)

func LoginFilter(ctx *context.Context) {
	// &符号的意思是对变量取地址,如：变量a的地址是&a
	// *符号的意思是对指针取值,如:*&a,就是a变量所在地址的值,当然也就是a的值了
	// *和 & 可以互相抵消,同时注意,*&可以抵消掉,但&*是不可以抵消的
	// a和*&a是一样的,都是a的值,值为1 (因为*&互相抵消掉了)
	// 同理,a和*&*&*&*&a是一样的,都是1 (因为4个*&互相抵消掉了)
	RequestUri := ctx.Request.RequestURI
	_, ok := (*GetLoginWhiteList())[RequestUri]
	if !ok {
		_, ok := ctx.Input.Session("UserName").(string)
		if !ok {
			ctx.Redirect(302, "/user/login")
		}
	}
}

var LoginWhiteList *map[string]string				// 登录白名单
var once sync.Once

func GetLoginWhiteList() *map[string]string {
	once.Do(func() {
		m := make(map[string]string)
		m["/user/login"] = "/user/login"
		m["/user/login/"] = "/user/login/"
		m["/user/regist"] = "/user/regist"
		m["/user/regist/"] = "/user/regist/"
		LoginWhiteList = &m
	})
	return LoginWhiteList
}

